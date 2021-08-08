#!/bin/bash
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

source $SCRIPT_DIR/config.sh 
NOW=$(date +"%m-%d-%Y")
MERGED404PATH=`find /var/lib/docker/overlay2/ -name '404.html' | grep merged `
SRC="$SCRIPT_DIR/../html/"
DEST="$(dirname ${MERGED404PATH})/"

#clear
rm -rf $SRC 
#untar recreate SRC
cd "$SCRIPT_DIR/.."
tar -xvf "$SCRIPT_DIR/../html.tar.gz"
echo "======= check source directory ========= "
ls -R $SRC

# command looks like
# "rsync --delete -avz /home/sp/goweb/html/
# /var/lib/docker/overlay2/511128e41cf1c2df2f270176d6f8de647ef5bcffe44e0d2e2317a4914123e03c/merged/app/html/
# >> /home/sp/goweb/rsync-html-$NOW.log"
echo "=====   rsync --delete -avz $SRC $DEST >> ${SCRIPT_DIR}/rsync-html-$NOW.log  ======"
rsync --delete -avz $SRC $DEST >> ${SCRIPT_DIR}/rsync-html-$NOW.log
echo "--- remote log for rsync ---- "
cat ${SCRIPT_DIR}/rsync-html-$NOW.log

