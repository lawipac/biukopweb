#/bin/bash
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
PROJ_DIR="$SCRIPT_DIR/../"

echo "build binary "
go build -o goweb

echo "change to $PROJ_DIR .."
cd $PROJ_DIR

echo "update web $PROJ_DIR/deploy/biukopweb-html/"
cd $PROJ_DIR/deploy/biukopweb-html/
git pull

echo "enter $PROJ_DIR"
cd $PROJ_DIR
echo "build docker image"
docker build -t biukopweb:1.0 $PROJ_DIR
docker image tag biukopweb:1.0 lawipac/biukopweb:1.0
docker image tag biukopweb:1.0 lawipac/biukopweb:latest

echo "list docker image"
docker image ls | grep biukopweb

echo "publish to docker hub"
docker push lawipac/biukopweb:latest



