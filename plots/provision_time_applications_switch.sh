
# Just add a new variable for an application 
# run the application

MEASUREMENTS=10
COUNT=1

MANIFEST_FOLDER=/home/master/p4kube/manifests
APPLICATION_FOLDER=/home/master/p4kube/Applications

APPLICATION=$1
NETWORK_FUNCTION=$2
YAML_APP=$3
YAML_NF=$4

DESIRED_APPLICATION_STATUS="Running"
DESIRED_NF_STATUS="Deployed"

total_time=0

while [ $COUNT -le $MEASUREMENTS ]
do
  start_time=$(date +%s%N)
  
  kubectl create -f $APPLICATION_FOLDER/$YAML_APP
  kubectl create -f $MANIFEST_FOLDER/$NETWORK_FUNCTION/$YAML_NF

  while true; 
  do
    status_nf=$(kubectl get p4s -n p4-namespace -o jsonpath='{.items[2].status.progress}')
    status_app=$(kubectl get pods -o jsonpath='{.items[0].status.phase}')

    if [ "$status_nf" == "$DESIRED_NF_STATUS" ] && [ "$status_app" == "$DESIRED_APPLICATION_STATUS" ]; then
      break
    fi
  done

  prov_time=$((($(date +%s%N) - $start_time)/1000000))
  result=$(echo "scale=2; $prov_time / 1000" | bc)
  echo "Provisiong time : $result"

  kubectl delete -f $APPLICATION_FOLDER/$YAML_APP
  kubectl delete -f $MANIFEST_FOLDER/$NETWORK_FUNCTION/$YAML_NF
  
  ((COUNT++))
done
