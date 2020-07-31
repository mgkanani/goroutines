#!/usr/bin/env bash

cd synctest

for ((i=1;i<=100000;i=i*10)) ; do
    echo "entries = $i"
    go test -bench=SyncMapLoad -entries=${i} -sync=true | grep SyncMapLoad
    go test -bench=StdMapLoad -entries=${i} -sync=false | grep StdMapLoad
done
