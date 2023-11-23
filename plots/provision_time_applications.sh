
# Just add a new variable for an application 
# run the application

MEASUREMENTS=50
COUNT=1

MANIFEST_FOLDER=/home/master/p4kube/manifests
APPLICATION_FOLDER=/home/master/p4kube/Applications
PLOT_FOLDER=/home/master/p4kube/plots

APPLICATION=$1
NETWORK_FUNCTION=$2
YAML_APP=$3
YAML_NF=$4

DESIRED_APPLICATION_STATUS="Running"
DESIRED_NF_STATUS="Deployed"


while [ $COUNT -le $MEASUREMENTS ]
do
  echo "Measurement: $COUNT"
  ./testbed_setting.sh $NETWORK_FUNCTION
  start_time=$(date +%s%N)

  kubectl create -f $APPLICATION_FOLDER/$YAML_APP
  kubectl create -f $MANIFEST_FOLDER/$NETWORK_FUNCTION/$YAML_NF

  while true; 
  do
    status_nf=$(kubectl get p4s -n p4-namespace --sort-by=.metadata.creationTimestamp -o jsonpath='{.items[-1:].status.progress}')
    status_app=$(kubectl get pods --sort-by=.metadata.creationTimestamp -o jsonpath='{.items[-1:].status.phase}')
    if [ "$status_nf" == "$DESIRED_NF_STATUS" ] && [ "$status_app" == "$DESIRED_APPLICATION_STATUS" ]; then
      break
    fi
  done

  prov_time=$((($(date +%s%N) - $start_time)/1000000))
  result=$(echo "scale=2; $prov_time / 1000" | bc)
  echo "The provisionig time: $result seconds"
  
  echo $result >> file1.txt

  source $PLOT_FOLDER/clean.sh

  ((COUNT++))
done
