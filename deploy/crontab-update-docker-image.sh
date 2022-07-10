#!/bin/bash

#crontab example setup
#0 5 * * * /root/update_biukopweb_docker.sh >> /root/update_docker.log

DATE=$(date)
IMAGE="lawipac/biukopweb"
echo "$DATE" Puling $IMAGE ...
out=$(docker pull $IMAGE)
echo "$out"

#there is real update from docker pull
if [[ $out != *"up to date"* ]]; then
        echo NEW image pulled on "$DATE"
        docker stop biukopweb-container
        docker rm   biukopweb-container
        docker run --name=biukopweb-container --restart=unless-stopped -p 8080:8080 -d $IMAGE
        docker image prune -f
else
        echo "$DATE" - docker pull has no new image
fi
