
# Scheduling Usecase 3
# The p42c and c-compilation is already executed
# usage: ./provision_time_usecase3.sh firewall firewall-switch.yaml

MEASUREMENTS=10
COUNT=1

MANIFEST_FOLDER=/home/master/p4kube/manifests
NETWORK_FUNCTION=$1
STATUS=""
DESIRED_STATUS="Deployed"
total_time=0
YAML=$2

while [ $COUNT -le $MEASUREMENTS ]
do
  cd $MANIFEST_FOLDER/$NETWORK_FUNCTION
  start_time=$(date +%s%N)

  kubectl create -f $YAML

  while true; 
  do
    status=$(kubectl get p4s -n p4-namespace -o jsonpath='{.items[2].status.progress}')
    if [ "$status" == "$DESIRED_STATUS" ]; then
      break
    fi
  done

  prov_time=$((($(date +%s%N) - $start_time)/1000000))
  total_time=$(($total_time+$prov_time))

  kubectl delete -f $YAML

  ((COUNT++))
done
average_time=$(($total_time/$MEASUREMENTS))
result=$(echo "scale=2; $average_time / 1000" | bc)
echo "The average provisionig time: $result seconds"
