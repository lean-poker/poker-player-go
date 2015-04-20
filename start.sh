#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

cd $DIR
go build player-service.go

PORT=4711 $DIR/player-service &
echo $! > $DIR/go.pid
