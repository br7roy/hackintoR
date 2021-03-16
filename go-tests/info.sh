#!/bin/env bash
base_json=`curl -s --negotiate -u : "bd15-21-33-61:10880/ws/v1/cluster/scheduler"`
queue_json=`echo $base_json|jq '.scheduler.schedulerInfo.rootQueue'`
queue_json_size=`echo $queue_json|jq '.childQueues[]|.queueName'|wc -l`
res=""
for((j=0;j<$queue_json_size;j++))
do
queue_name=`echo $queue_json|jq ".childQueues[$j].queueName"|sed 's/\"//g'`
res=$queue_name","$res
done
echo $res
