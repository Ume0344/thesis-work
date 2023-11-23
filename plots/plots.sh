echo "Collecting measurements for port-fwd" >> file.txt
echo "running usecase1" >> file.txt
./provision_time.sh portfwd-gen portfwd-gen-target-node.yaml 

echo "running usecase2" >> file.txt
./provision_time.sh portfwd-gen portfwd-gen-random.yaml

echo "running usecase3" >> file.txt
./provision_time.sh portfwd-gen portfwd-gen-switch.yaml

#######################################################################################################
echo "Collecting measurements for firewall-l2fwd" >> file.txt
echo "running usecase1" >> file.txt
./provision_time_combination.sh firewall l2fwd-gen firewall-target-node.yaml l2fwd-gen-target-node.yaml

echo "running usecase2" >> file.txt
./provision_time_combination.sh firewall l2fwd-gen firewall-random.yaml l2fwd-gen-random.yaml

echo "running usecase3" >> file.txt
./provision_time_combination.sh firewall l2fwd-gen firewall-switch.yaml l2fwd-gen-switch.yaml

#######################################################################################################
echo "Collecting measurements for firewall-portfwd" >> file.txt
echo "running usecase1" >> file.txt
./provision_time_combination.sh firewall portfwd-gen firewall-target-node.yaml portfwd-gen-target-node.yaml

echo "running usecase2" >> file.txt
./provision_time_combination.sh firewall portfwd-gen firewall-random.yaml portfwd-gen-random.yaml

echo "running usecase3" >> file.txt
./provision_time_combination.sh firewall portfwd-gen firewall-switch.yaml portfwd-gen-switch.yaml

#######################################################################################################
echo "Collecting measurements for l2fwd-portfwd" >> file.txt
echo "running usecase1" >> file.txt
./provision_time_combination.sh l2fwd-gen portfwd-gen l2fwd-gen-target-node.yaml portfwd-gen-target-node.yaml

echo "running usecase2" >> file.txt
./provision_time_combination.sh l2fwd-gen portfwd-gen l2fwd-gen-random.yaml portfwd-gen-random.yaml

echo "running usecase3" >> file.txt
./provision_time_combination.sh l2fwd-gen portfwd-gen l2fwd-gen-switch.yaml portfwd-gen-switch.yaml
