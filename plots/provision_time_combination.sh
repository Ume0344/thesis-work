
# Scheduling
# run the firewall random, targeted and splitted and find the time
# run it for 10 times

MEASUREMENTS=10
COUNT=1

MANIFEST_FOLDER=/home/master/p4kube/manifests
NETWORK_FUNCTION_0=$1
NETWORK_FUNCTION_1=$2
STATUS=""
DESIRED_STATUS="Deployed"
total_time=0
YAML_0=$3
YAML_1=$4

start_time=$(date +%s%N)
while [ $COUNT -le $MEASUREMENTS ]
do
  for ((i=0; i<=1; i++)); do
    network_function="NETWORK_FUNCTION_$i"
    yaml="YAML_$i"
    kubectl create -f $MANIFEST_FOLDER/${!network_function}/${!yaml}
    while true; 
    do
      status=$(kubectl get p4s -n p4-namespace -o jsonpath='{.items['$i'].status.progress}')
      if [ "$status" == "$DESIRED_STATUS" ]; then
          break
      fi
    done
  done

  ((COUNT++))
done

prov_time=$((($(date +%s%N) - $start_time)/1000000))
result=$(echo "scale=2; $prov_time / 1000" | bc)
echo "Provisiong time : $result"
