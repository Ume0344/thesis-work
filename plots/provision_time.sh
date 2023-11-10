
# Scheduling
# run the firewall random, targeted and splitted and find the time
# run it for 10 times

MEASUREMENTS=10
COUNT=1

MANIFEST_FOLDER=/home/master/p4kube/manifests
NETWORK_FUNCTION=$1
STATUS=""
DESIRED_STATUS="Deployed"
total_time=0
YAML=$2

workernodes=("p4-kube@192.168.0.211" "p4-kube@192.168.0.186" "p4-kube@192.168.0.199")

while [ $COUNT -le $MEASUREMENTS ]
do
  cd $MANIFEST_FOLDER/$NETWORK_FUNCTION
  start_time=$(date +%s%N)

  kubectl create -f $YAML

  while true; 
  do
    status=$(kubectl get p4s -n p4-namespace -o jsonpath='{.items[0].status.progress}')
    if [ "$status" == "$DESIRED_STATUS" ]; then
      break
    fi
  done

  prov_time=$((($(date +%s%N) - $start_time)/1000000))
  total_time=$(($total_time+$prov_time))

  kubectl delete -f $YAML
  
  echo "Deleting build files on nodes..."
  for node in "${workernodes[@]}";
  do
    ssh $node "sudo rm -rf /home/p4-kube/t4p4s/build/last /home/p4-kube/t4p4s/build/aftermost /home/p4-kube/t4p4s/build/$NETWORK_FUNCTION@std-v1model"
  done
  echo "Build files deleted on nodes"

  ((COUNT++))
done
average_time=$(($total_time/$MEASUREMENTS))
result=$(echo "scale=2; $average_time / 1000" | bc)
echo "The average provisionig time: $result seconds"
