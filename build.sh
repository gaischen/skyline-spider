#! /bin/bash

##处理下网络
env|grep -I proxy
git config --global http.proxy

##更新代码
git pull
go mod vendor

##编译
echo "start building main.go..."
go build main.go
echo "build main.go successful"

##进程历史
pid=$(ps -ef | grep "main" | grep -v grep | awk '{print $2}')
echo "found go pid:$pid"

if ps -p $pid > /dev/null
  then
     echo "$pid is running"
     echo "killing $pid"
     kill -9 $pid
      echo "killing $pid successful"
fi


nohup  ./main > run.log 2>&1 &

npid=$(ps -ef | grep "main" | grep -v grep | awk '{print $2}')
echo "start new process...new pid is: $npid"