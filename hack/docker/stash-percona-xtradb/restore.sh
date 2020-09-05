#!/usr/bin/env bash

# Copyright AppsCode Inc. and Contributors
#
# Licensed under the AppsCode Community License 1.0.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

datadir=$1

xbstream -x -C $datadir
xtrabackup --prepare --target-dir=$datadir

# give the ownership (mysql:mysql) needed
groupadd -g 1001 mysql
useradd -u 1001 -r -g 1001 -s /sbin/nologin \
    -c "Default Application User" mysql
chown -R mysql:mysql $datadir
