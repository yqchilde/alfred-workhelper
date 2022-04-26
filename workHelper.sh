#!/bin/bash

if [[ $1 == "build" ]]; then
  go build -ldflags="-s -w" -o workHelper main.go
  #  Apple Silicon upx待兼容
  # upx workHelper
elif [[ $1 == "run" ]]; then
  export alfred_workflow_bundleid="net.deanishe.awgo"
  export alfred_workflow_data="./data"
  export alfred_workflow_cache="./cache"
  # go run main.go arg1 arg2 ...
  # examples: go run main.go date now
fi