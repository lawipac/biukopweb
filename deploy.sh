#/bin/bash

rm -rf /tmp/goweb
mkdir -p /tmp/goweb
cp -a . /tmp/goweb/
rm -rf /tmp/goweb/html/*
rsync -avh /mnt/hgfs/workspace/2021-07-31-BiukopWeb/ /tmp/goweb/html/

cd /tmp/goweb
gcloud app deploy