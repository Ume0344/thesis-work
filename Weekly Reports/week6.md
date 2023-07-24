1. Download DPDK, dpdk-22.03, untar it `tar xJf dpdk-22.03.tar.xz dpdk-22.03/`.
    Following this [tutorial](https://doc.dpdk.org/guides/linux_gsg/build_dpdk.html) 
2. Downaload meson; `sudo pip install meson==0.59.0`
3. Install elftools; `sudo pip install pyelftools`.
4. go to dpdk directory; Run; `meson setup build`
    go to build/ and run `ninja`
5. Run; `sudo meson install` 
    got the error; `AttributeError: Can't get attribute 'OctalInt' on <module 'mesonbuild.coredata' from '/usr/lib/python3/dist-packages/mesonbuild/coredata.py'>`. 
    Resoolved it by installing meson with root access and runnin the command again.
6. Run; sudo ldconfig

7. To confirm if installation is done correctly, Run hello world program in the examples by; `cd examples/helloworld && make && cd ./build` and `sudo ./helloworld –l 0-3 –n 2` ; it will print helloworld. 


Some Concepts: 
**VFIO-PCI** : VFIO-PCI is a Linux kernel driver and framework that provides direct access to PCI devices for userspace applications.
**Poll Mode Drivers** :  Software modules or drivers that provide an interface between the DPDK framework and the underlying network devices or hardware.
**Kernel drivers** : also known as device drivers or simply drivers, are software components that enable communication and interaction between hardware devices and the operating system's kernel. Mostly, we need to unbound the devices using dpdk from kernel and bound it to vfio-pci kernel module before application is run.
**HugePages** :  hugepages refer to a memory management technique that enables the allocation of large, contiguous memory blocks for improved performance and reduced memory management overhead. This help to rapidly process the network packets.Typically, operating systems allocate memory in smaller page sizes, typically 4 KB (a bottle neck to high procesing). Hugepages allow dpdk to have large page sizes , 2MB or more. 

**Some Questions** : how dpdk bind kernel deriver module vfio-pci to network devices?
what is the impact of large hugepages on processing speed
how can i deploy a network function on nic with dpdk and get the confirmation that output is right.
I need to watch some tutorials on it. 

**To look into what NIC are available** :
I can run `lspci` to find about the available nic. If I do, `lspci | grep Ethernet`, it will show the thernet connections. 

**Common Commands**
`./dpdk-devbind.py -s`
`cd /sys/bus/pci/drivers` to see drivers
`sudo lspci`

Setting up hugepage; `sudo dpdk-hugepages.py -p 2M --setup 20M`

First of all load vfci-pci module; I cannot load it, so I used `ufio_pci_generic` driver.
    `lsmod | grep <drivername >` # to see if driver is loaded or not
    `sudo modprobe < driver_name>` # to load driver

To bind a PCI device, `sudo ./dpdk-devbind.py -b uio_pci_generic enp4s0`
To unbind, `sudo ./dpdk-devbind.py -u <interface>`

To run t4p4s for l2fwd, go to t4p4s directory, `./t4p4s.sh ./examples/p4_14/l2fwd.p4_14 model=v1model`
`./t4p4s.sh :l2fwd model=v1model cores=2 ports=1 verbose=1`


**Running l2fwd with only dpdk**
`sudo dpdk-hugepages.py -p 2M --setup 50M` and `sudo ./l2fwd -l 0-1 -- -q 1 -p 1`

**Some more questions**
- How to know information about available NIC on system/server.
  by running `lspci`, it will list all the available network interfaces (L2, L3)


  **What is P4-DPDK target**
  https://www.youtube.com/watch?v=dPvH_joaScA

  A tool to run p4 programs on multi-core cpu using dpdk.
  - workflow : P4 program -> P4 compiler -> creates .spec file -> dpdk library -> creates c program based on .spec fil
  e -> C compiler -> a binary representation to be deployed on NIC.

  https://www.youtube.com/watch?v=xJR-5DcqhlY

  **P4 Runtime**





















































