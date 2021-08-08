#!/bin/bash
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
source $SCRIPT_DIR/gcp_version.sh

#wait GAE pending operations
for ((i=1; i<120; i++)); do
  echo "check pending GAE operations ... "
  op=`gcloud app operations list --filter=status=PENDING --format=json`
  if [ "$op" == "[]" ]; then 
    echo "no pending"; 
    del_old_instances
    break; 
  else
    echo "$op"
    echo "check $i of 120  (normally it takes 70 checks = 3-5 minutes"
    sleep 3;
  fi 
done
echo "finished";
