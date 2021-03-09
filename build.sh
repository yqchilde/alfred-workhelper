#!/bin/bash

for file in $(ls cmd)
do
  go build "./cmd/$file"
  upx "./$file"
done