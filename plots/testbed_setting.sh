# 1 empty node
# 1 p4-c node
# 1 c-compilation node 

MANIFEST_FOLDER=/home/master/p4kube/manifests

NETWORK_FUNCTION=$1
YAML=("$NETWORK_FUNCTION-p42c.yaml" "$NETWORK_FUNCTION-c-compile.yaml")

for ((i=0; i<3; i++));
do
  random_number=$((RANDOM % 3))
  echo $random_number
  if [ $random_number == 2 ];then
    continue
  fi
  
  file=${YAML["$random_number"]}
  echo "Manifest file to be deployed: $file"
  kubectl create -f $MANIFEST_FOLDER/$NETWORK_FUNCTION/$file
  sleep 5
done
