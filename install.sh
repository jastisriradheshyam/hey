#!/bin/bash

cd go
go build -ldflags "-s -w" -trimpath
chmod 775 hey
sudo cp hey /bin
