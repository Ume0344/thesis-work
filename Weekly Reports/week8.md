
**T4P4S and P4Runtime**

We will use T4P4S for the thesis.
T4P4S does not use P4Runtime. We can integrate P4Runtime to T4P4S. One of the proposal could be; 

t4p4s compiles and deploys the P4 program to NIC and also generates p4info.txt file. As p4c along with backend compiler always generates binary and P4info.txt file, we can use p4info.txt file for p4runtime. But in t4p4s it is generating HLIR in the form of json (however in our case, it should produce HLIR and p4info.txt). We can look into t4p4s source code and see how it is compiling. We can change it to produce p4info.txt along with HLIR.json file to be used by p4runtime. After we get p4info.txt, we can use p4runtime shell environment to change the deployed p4 program. 


t4p4s is using [hlir16](https://github.com/P4ELTE/hlir16/tree/8ee85b84478567658ec48fae7c3e0357490cfa86) which uses p4c to generate json file from p4 program. In this .py file, it can be seen that hlir is using p4c to generate [json file](https://github.com/P4ELTE/hlir16/blob/8ee85b84478567658ec48fae7c3e0357490cfa86/hlir.py#L103). 

https://opennetworking.org/wp-content/uploads/2019/10/NG-SDN-Tutorial-Session-1.pdf Slide 52 shows how p4 program is deployed with p4c


**Helpful Commands**
To convert compiler IR of P4 program into json format.
/home/p4kube/p4c/build/p4test ab.p4 --toJSON ab.json 

(ab.p4 and ab.json are test files)
