#/bin/bash
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
PROJ_DIR="$SCRIPT_DIR/../"

rm -rf /tmp/goweb
echo "create /tmp/goweb"
mkdir -p /tmp/goweb
echo "prepare golang project"
rsync  -a  $PROJ_DIR /tmp/goweb/
rm -rf /tmp/goweb/html/*
echo "sync Web html"
rsync -a /mnt/hgfs/workspace/2021-07-31-BiukopWeb/ /tmp/goweb/html/

cd /tmp/goweb
gcloud app deploy

#list all versions
gcloud app versions list ;



