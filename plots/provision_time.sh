
# Scheduling
# run the firewall random, targeted and splitted and find the time
# run it for 10 times

MEASUREMENTS=10
COUNT=1

MANIFEST_FOLDER=/home/master/p4kube/manifests
PLOT_FOLDER=/home/master/p4kube/plots

NETWORK_FUNCTION=$1
STATUS=""
DESIRED_STATUS="Deployed"
total_time=0
YAML=$2

while [ $COUNT -le $MEASUREMENTS ]
do
  ./testbed_setting.sh $NETWORK_FUNCTION
  start_time=$(date +%s%N)

  kubectl create -f $MANIFEST_FOLDER/$NETWORK_FUNCTION/$YAML

  while true; 
  do
    status=$(kubectl get p4s -n p4-namespace -o jsonpath='{.items[0].status.progress}')
    if [ "$status" == "$DESIRED_STATUS" ]; then
      break
    fi
  done

  prov_time=$((($(date +%s%N) - $start_time)/1000000))
  result=$(echo "scale=2; $prov_time / 1000" | bc)
  echo "The provisionig time: $result seconds"

  source $PLOT_FOLDER/clean.sh

  ((COUNT++))
done
