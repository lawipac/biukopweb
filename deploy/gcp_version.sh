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
  get_instance_id
  echo "RUNNING VERSION: " $LAST_VERSION_ID

  GCP_VERSIONS=(` gcloud app versions list | awk -F' ' '{print $2}' `)
  LENGTH=${#GCP_VERSIONS[@]}
  for i in $(seq 1 1 $(expr $LENGTH - 1) )
  do
    if [ "${GCP_VERSIONS[$i]}" != "$LAST_VERSION_ID" ] ; then
      echo "delete old version ${GCP_VERSIONS[$i]} "
      del_instance_by_version ${GCP_VERSIONS[$i]}
    else
      echo "KEEP RUNNING VERSION ${GCP_VERSIONS[$i]}"
    fi
  done
}

del_instance_by_version() {
  versionId=$1
  if [ "$versionId" == "" ]; then
    return
  fi
  gcloud app versions delete --service=default $versionId ;

#  while true; do
#      read -p "Do you wish to delete version $versionId ?" yn
#      case $yn in
#          [Yy]* )
#            echo gcloud app instances delete $INSTANCE_ID --service=default --version=$versionId ;
#            break;;
#          [Nn]* ) exit;;
#          * ) echo "Please answer yes or no.";;
#      esac
#  done

}

