<center><h2><b>Week 14: 04.09 - 08.09.23</b></h2></center>

- This week I looked into source code of T4P4s. After, going through it, I found that t4p4s.sh script is the place to divide the t4p4s whole process. From t4p4s.sh script I got to know that the t4p4s process is already divided into 3 process;
    - P4 to C conversion
    - C Compilation
    - Running the switch

- I deployed these as separate custom resources on kubernetes cluster. The average provisioning / compilation time for l2fwd-gen.p4 program is;
    - P4 to C conversion - 2.7s (Here, the error was comming initially. After debugging, I got to know that the exit status is an empty value and controller detects it and deployment is unsuccessful. I need more time to dig into it to figure out why exit status is empty instead of 0.)
    - C Compilation - 2.6s
    - Running the switch - 2.5s
