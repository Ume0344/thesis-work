
# Scheduling
# run the firewall random, targeted and splitted and find the time
# run it for 10 times

MEASUREMENTS=50
COUNT=1

MANIFEST_FOLDER=/home/master/p4kube/manifests
PLOT_FOLDER=/home/master/p4kube/plots

NETWORK_FUNCTION_0=$1
NETWORK_FUNCTION_1=$2

STATUS=""
DESIRED_STATUS="Deployed"

YAML_0=$3
YAML_1=$4

while [ $COUNT -le $MEASUREMENTS ]
do
  echo "Measurement: $COUNT"
  ./testbed_setting.sh $NETWORK_FUNCTION_0
  ./testbed_setting.sh $NETWORK_FUNCTION_1
  start_time=$(date +%s%N)
  for ((i=0; i<=1; i++)); do
    network_function="NETWORK_FUNCTION_$i"
    yaml="YAML_$i"
    kubectl create -f $MANIFEST_FOLDER/${!network_function}/${!yaml}
    while true; 
    do
      status=$(kubectl get p4s -n p4-namespace --sort-by=.metadata.creationTimestamp -o jsonpath='{.items[-1:].status.progress}')
      if [ "$status" == "$DESIRED_STATUS" ]; then
          break
      fi
    done

  done
  prov_time=$((($(date +%s%N) - $start_time)/1000000))
  result=$(echo "scale=2; $prov_time / 1000" | bc)
  echo "The provisionig time: $result seconds"
  echo $result >> file.txt

  source $PLOT_FOLDER/clean.sh

  ((COUNT++))
done