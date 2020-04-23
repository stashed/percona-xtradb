# stash-percona-xtradb

[stash-percona-xtradb](https://github.com/stashed/percona-xtradb) - Percona XtraDB database backup/restore plugin for [Stash by AppsCode](https://appscode.com/products/stash/).

## TL;DR;

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm install stash-percona-xtradb-5.7 appscode/stash-percona-xtradb -n kube-system --version=5.7
```

## Introduction

This chart installs necessary `Functions` and `Tasks` definitions to take backup of Percona XtraDB 5.7 databases and restore them using Stash.

## Prerequisites

- Kubernetes 1.11+

## Installing the Chart

- Add AppsCode chart repository to your helm repository list,

```console
$ helm repo add appscode https://charts.appscode.com/stable/
```

- Update helm repositories to fetch latest charts from the remove repository,

```console
$ helm repo update
```

- Install the chart with the release name `stash-percona-xtradb-5.7` run the following command,

```console
$ helm install stash-percona-xtradb-5.7 appscode/stash-percona-xtradb -n kube-system --version=5.7
```

The above commands installs `Functions` and `Task` CRDs that are necessary to take backup of Percona XtraDB 5.7 databases and restore them using Stash.

## Uninstalling the Chart

To uninstall/delete the `stash-percona-xtradb-5.7` run the following command,

```console
helm uninstall stash-percona-xtradb-5.7 -n kube-system --purge
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `stash-percona-xtradb` chart and their default values.

|     Parameter        |                                                                    Description                                                                     |      Default      |
| :------------------: | -------------------------------------------------------------------------------------------------------------------------------------------------- | :---------------: |
| `image.registry`     | Docker registry used to pull respective images                                                                                                     |     `stashed`     |
| `image.repository`   | Docker image used to take backup of Percona XtraDB databases and restore them                                                                               |   `stash-percona-xtradb`   |
| `image.tag`          | Tag of the image that is used to take backup of Percona XtraDB databases and restore them. This is usually same as the database version it can take backup. |       `5.7`    |
| `backup.args`  | Optional arguments to pass to `mysqldump` command  during bakcup process                                                                           | `--all-databases` |
| `restore.args` | Optional arguments to pass to `mysql` command during restore process                                                                               |        ""         |

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`.

For example:

```console
helm install stash-percona-xtradb-5.7 appscode/stash-percona-xtradb -n kube-system ---set image.registry=my-registry
```
