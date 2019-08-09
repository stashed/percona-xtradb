#!/usr/bin/env bash

datadir=$1

xbstream -x -C $datadir
xtrabackup --prepare --target-dir=$datadir

# give the ownership (mysql:mysql) needed
groupadd -g 1001 mysql
useradd -u 1001 -r -g 1001 -s /sbin/nologin \
            -c "Default Application User" mysql
chown -R mysql:mysql $datadir