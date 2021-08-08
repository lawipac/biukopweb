#/bin/bash
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
PROJ_DIR="$SCRIPT_DIR/../"

rm -rf /tmp/goweb
mkdir -p /tmp/goweb
rsync  -a  $PROJ_DIR /tmp/goweb/
rm -rf /tmp/goweb/html/*
rsync -avh /mnt/hgfs/workspace/2021-07-31-BiukopWeb/ /tmp/goweb/html/

cd /tmp/goweb
#gcloud app deploy

source $SCRIPT_DIR/gcp_version.sh
del_old_instances