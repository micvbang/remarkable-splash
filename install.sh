#!/usr/bin/env bash

BASE=$(pwd)

cd ./cmd/remarkable-splash && \
    env GOOS=linux GOARCH=arm GOARM=7 go build -o remarkable-splash && \
    scp remarkable-splash remarkable:

cd $BASE
scp ./xkcd.service remarkable:/etc/systemd/system/
ssh remarkable "systemctl daemon-reload && systemctl enable xkcd && systemctl restart xkcd"
