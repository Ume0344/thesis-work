workernodes=("p4-kube@192.168.0.211" "p4-kube@192.168.0.186" "p4-kube@192.168.0.199")

echo "Deleting build files on nodes..."
for node in "${workernodes[@]}";
do
  ssh $node "sudo rm -rf /home/p4-kube/t4p4s/build/last /home/p4-kube/t4p4s/build/aftermost \
  /home/p4-kube/t4p4s/build/l2fwd-gen@std-v1model /home/p4-kube/t4p4s/build/portfwd-gen@std-v1model \
  /home/p4-kube/t4p4s/build/firewall@std-v1model"
done
echo "Build files deleted on nodes"

echo "Deleting p4 resources"
kubectl delete p4s -n p4-namespace --all

echo "Deleting applications"
kubectl delete deployments --all
kubectl delete svc --all

echo "Done"
