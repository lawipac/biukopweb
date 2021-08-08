#!/bin/bash

source config.sh
source gcp_version.sh

get_instance_id
echo INSTANCE ID: $INSTANCE_ID
echo VERSION: $LAST_VERSION_ID

rm    -rf $HTML_WORKDIR
mkdir -p  $HTML_WORKDIR

rsync -a $PROJ_DIR/deploy $HTML_WORKDIR
mkdir -p "$HTML_WORKDIR/html"
rsync -a /mnt/hgfs/workspace/2021-07-31-BiukopWeb/* "$HTML_WORKDIR/html/"
#create compressed archive
echo tar -zcvf -C $HTML_WORKDIR $HTML_WORKDIR/html.tar.gz "$HTML_WORKDIR/html/"
tar -C $HTML_WORKDIR -zcvf $HTML_WORKDIR/html.tar.gz   "html"
rm -rf "$HTML_WORKDIR/html/"

rm -f "/home/sp/.ssh/google_compute_known_hosts"

# copy deploy script and html to destination
echo "===copy $HTML_WORKDIR to remote instance ======"
gcloud app instances scp --recurse "$HTML_WORKDIR"  $INSTANCE_ID:/home/sp/  --service=default --version=$LAST_VERSION_ID
# copy html
echo "===execute remote sync======"
gcloud app instances ssh $INSTANCE_ID --service=default --version=$LAST_VERSION_ID -- sudo bash $REMOTE_WORK/deploy/r_sync_html.sh
echo "=== update complete this put the instance into debug mode == "
gcloud app instances disable-debug
