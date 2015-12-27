#!/bin/bash

host=i3@192.168.1.245
des=/home/i3/code/go_code/src/github.com/beego/samples
ssh $host "rm -rf $des/*"
scp -r * $host:$des
