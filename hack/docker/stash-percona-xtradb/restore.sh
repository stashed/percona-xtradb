#!/usr/bin/env bash

# Copyright The Stash Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
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