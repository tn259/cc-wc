#!/bin/bash
set -o errexit

trap 'echo "Error: $0:$LINENO failed with exit code $?" >&2' ERR

go run main.go -c test.txt | grep -q '342190'
go run main.go -l test.txt | grep -q '7145'
go run main.go -w test.txt | grep -q '58164'
go run main.go -m test.txt | grep -q '339292'
go run main.go test.txt | grep -q '342190 7145 58164 339292'

cat test.txt | go run main.go -l | grep -q '7145'
echo 'PASSED'