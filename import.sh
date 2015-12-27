#!/bin/bash

host=i3@192.168.1.245
src=/home/i3/code/go_code/src/github.com/beego/samples/todo
scp -r $host:$src/* ./todo
