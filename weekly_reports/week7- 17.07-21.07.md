 <center><h2><b>Week 7: 17.07 - 21.07.23</b></h2></center>

- This week we surveyed P4Runtime and how it works. 
  
**P4 Runtime**

  P4 Runtime is a way for control plane software to control the forwarding plane of a switch, router, firewall, load-balancer, etc. Perhaps the most novel aspect of P4 Runtime is that it lets you control any forwarding plane, regardless of whether it is built from a fixed-function or programmable switch ASIC, an FPGA, NPU or a software switch running on an x86 server. 

  **How does it work**?

  We write p4 program, a compiler compiles P4 program into protobuf based file  .p4info to have all the rules (add/drop an entry etc). This protobuf-based file is used by p4runtime api to control the switch dataplane.

  **Running a ping example by using P4Rutnime Shell Environemnt**

 [Tutorial](https://github.com/opennetworkinglab/ngsdn-tutorial/blob/advanced/EXERCISE-1.md) 

 Here, we run p4 program, compile it with p4c. This produces p4info file based on protobuf format. Started p4runtime shell environment. This p4runtime environment uses p4info file to desrcibe tables, entries etc.

  Mostly, we can follow the tutorial as it is till this section *3. Start Mininet topology*. In section *Program leaf1 using P4Runtime*, follow the below commands;
  - Enter into P4Runtime Shell; 

    `util/p4rt-sh --grpc-addr localhost:50001 --config p4src/build/p4info.txt,p4src/build/bmv2.json --election-id 0,1`

  - In seperate tab, run mininet by `make mn-cli`.
  - Insert static NDP entries;

    `h1a ip -6 neigh replace 2001:1:1::B lladdr 00:00:00:00:00:1B dev h1a-eth0`

    `h1b ip -6 neigh replace 2001:1:1::A lladdr 00:00:00:00:00:1A dev h1b-eth0`

  - ping between h1a and h1b; `h1a ping h1b`

  - Add table entries in P4Runtime shell by using below commands;

    `global_options["canonical_bytestrings"] = False`

    `te = table_entry["IngressPipeImpl.l2_exact_table"](action = "IngressPipeImpl.set_egress_port")`

    `te.match["hdr.ethernet.dst_addr"] = ("00:00:00:00:00:1B")`
    `te.action["port_num"] = ("4")`

    `print(te)`

    `te.insert()`

    `te.match["hdr.ethernet.dst_addr"] = ("00:00:00:00:00:1A")`

    `te.action["port_num"] = ("3")`

    `print(te)`

    `te.insert()`

  - Switch back to mininet tab, now the two hosts are being able to ping.

  This tutorial also helped me; https://dr-kino.github.io/2021/03/30/next-generation-sdn-exercise-one/ 


 
 **How does T4P4S work**

T4P4S uses the official P4 compiler (P4C) to generate a JSON representation of the P4 program. This JSON object is then parsed by t4p4s using python, which in turn creates C code. This C code is platform independent and is linked with platform specific (DPDK) functions.

*While running l2fwd example in T4P4S,* At this step, it creates HLIR;

  ```[COMPILE  P4-14] ./examples/p4_14/l2fwd.p4_14 @std, debug mode, model v1model
P4 compiler options : --p4incdir=examples/include --p4opt=__TARGET_V1__ ./examples/p4_14/l2fwd.p4_14 --p4v 14 -g ./build/last/srcgen -verbose 
HLIR (cached: stage hlir_add_attributes) build/last/cache/l2fwd.p4_14.hlir.attributed.cached
```

And, we can see the HLIR in json format at this location;

`sudo nano /t4p4s/build/last/cache/l2fwd.p4_14.json.cached`
