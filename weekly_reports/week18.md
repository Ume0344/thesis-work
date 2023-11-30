<center><h2><b>Week 16: 02.10 - 06.10.23</b></h2></center>

- Helpful resources: https://banzaicloud.com/blog/k8s-custom-scheduler/ 

Steps to create custom scheduler

UseCase-1: Randomly schedules the p4resource on available nodes
- Watch out all the p4 resources which are not being scheduled.
- Find the available nodes.
- Bind the p4 resource randomly to one of the nodes.

UseCase-2: Schedules the p4resource on specified node (in manifest file)
- Watch out all the p4 resources which are not being scheduled.
- Check if specific node is available.
- If yes, bind the p4 resource to that node.

UseCase-3: Schedules the p4resource on a node that already has t4p4s build files.
- Watch out all the p4 resources which are not being scheduled.
- Find the node that has build files related to particular p4resource.
- Bind the p4 resource to that node.

Next week, I will start implementing the scheduling UseCase 1.
