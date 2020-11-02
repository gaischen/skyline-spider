#! /bin/bash

git config --global http.proxy

git pull

go mod vendor

ech "start building main.go..."

go build main.go

ech "build main.go successful"

pid = $(ps -ef | grep "main" | grep -v grep | awk '{print $2}')

echo "found go pid:" + $pid

if ps -p pid > /dev/null

then
   echo "$pid is running"
   # Do something knowing the pid exists, i.e. the process with $PID is running
fi


