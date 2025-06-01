#!/bin/bash

cd cmd/hey
go build -o ../../bin/hey -ldflags "-s -w" -trimpath
chmod 775 ../../bin/hey
sudo cp ../../bin/hey /bin
