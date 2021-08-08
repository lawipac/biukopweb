#!/bin/bash

GCP_VERSIONS=()
LAST_VERSION_ID=""
INSTANCE_ID=""

get_instance_id() {
#  SERVICE  VERSION          ID                                VM_STATUS  VM_LIVENESS  DEBUG_MODE
#  default  20210807t221459  aef-default-20210807t221459-1nqz  RUNNING    HEALTHY      YES
    echo "Getting gcloud instance id ..."
    if [ "$INSTANCE_ID" == "" ]
    then
      output=` gcloud app instances list | grep RUNNING`    # get the running version only
      INSTANCE_ID=(` echo $output | awk -F' ' '{print $3}' `)
      LAST_VERSION_ID=(` echo $output | awk -F' ' '{print $2}' `)
    fi
}

del_old_instances(){
  VERSION_OUTPUT=$(gcloud app versions list)
  # VERSION_OUTPUT=$` cat /home/sp/go/src/goweb/deploy/sample-version-list.txt`
  echo "$VERSION_OUTPUT"
  dropList=()

  while IFS= read -r line; do
    status=(` echo $line | awk -F' ' '{print $5}' `)
    version=(` echo $line | awk -F' ' '{print $2}' `)
    if [ "$status" == "STOPPED" ] ; then
       dropList+=("$version")
    fi
  done <<< "$VERSION_OUTPUT"

  #cannot do this within while, as the user input will by bypassed by <<<
  for version in  ${dropList[@]} ; do 
    gcloud app versions delete --service=default $version ;
  done
}
