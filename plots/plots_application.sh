#######################################################################################################
echo "Collecting measurements for nginx-firewall" >> file1.txt
echo "running usecase1" >> file1.txt
./provision_time_applications.sh nginx firewall nginx.yaml firewall-target-node.yaml

echo "running usecase2" >> file1.txt
./provision_time_applications.sh nginx firewall nginx.yaml firewall-random.yaml

echo "running usecase3" >> file1.txt
./provision_time_applications.sh nginx firewall nginx.yaml firewall-switch.yaml

#######################################################################################################
echo "Collecting measurements for wordpress-firewall" >> file1.txt
echo "running usecase1" >> file1.txt
./provision_time_applications.sh wordpress firewall wordpress.yaml firewall-target-node.yaml

echo "running usecase2" >> file1.txt
./provision_time_applications.sh wordpress firewall wordpress.yaml firewall-random.yaml

echo "running usecase3" >> file1.txt
./provision_time_applications.sh wordpress firewall wordpress.yaml firewall-switch.yaml

#######################################################################################################
echo "Collecting measurements for grafana-portfwd" >> file1.txt
echo "running usecase1" >> file1.txt
./provision_time_applications.sh grafana portfwd-gen grafana.yaml portfwd-gen-target-node.yaml

echo "running usecase2" >> file1.txt
./provision_time_applications.sh grafana portfwd-gen grafana.yaml portfwd-gen-random.yaml

echo "running usecase3" >> file1.txt
./provision_time_applications.sh grafana portfwd-gen grafana.yaml portfwd-gen-switch.yaml

#######################################################################################################
echo "Collecting measurements for nginx-l2fwd" >> file1.txt
echo "running usecase1" >> file1.txt
./provision_time_applications.sh nginx l2fwd-gen nginx.yaml l2fwd-gen-target-node.yaml

echo "running usecase2" >> file1.txt
./provision_time_applications.sh nginx l2fwd-gen nginx.yaml l2fwd-gen-random.yaml

echo "running usecase3" >> file1.txt
./provision_time_applications.sh nginx l2fwd-gen nginx.yaml l2fwd-gen-switch.yaml
