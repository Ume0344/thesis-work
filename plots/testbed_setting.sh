# 1 empty node
# 1 p4-c node
# 1 c-compilation node 

MANIFEST_FOLDER=/home/master/p4kube/manifests

NETWORK_FUNCTION=$1
YAML=("$NETWORK_FUNCTION-p42c.yaml" "$NETWORK_FUNCTION-p42c1.yaml" "$NETWORK_FUNCTION-c-compile.yaml")

for file in "${YAML[@]}";
do
  kubectl create -f $MANIFEST_FOLDER/$NETWORK_FUNCTION/$file
  sleep 3
done
