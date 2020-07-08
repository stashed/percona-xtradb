# stash-percona-xtradb

[stash-percona-xtradb](https://github.com/stashed/percona-xtradb) - Percona XtraDB database backup/restore plugin for [Stash by AppsCode](https://stash.run)

## TL;DR;

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm install stash-percona-xtradb-5.7-beta.20200708 appscode/stash-percona-xtradb -n kube-system --version=5.7-beta.20200708
```

## Introduction

This chart deploys necessary `Function` and `Task` definition to backup or restore Percona XtraDB 5.7 using Stash on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.11+

## Installing the Chart

To install the chart with the release name `stash-percona-xtradb-5.7-beta.20200708`:

```console
$ helm install stash-percona-xtradb-5.7-beta.20200708 appscode/stash-percona-xtradb -n kube-system --version=5.7-beta.20200708
```

The command deploys necessary `Function` and `Task` definition to backup or restore Percona XtraDB 5.7 using Stash on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `stash-percona-xtradb-5.7-beta.20200708`:

```console
$ helm delete stash-percona-xtradb-5.7-beta.20200708 -n kube-system
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `stash-percona-xtradb` chart and their default values.

|         Parameter         |                                                             Description                                                              |        Default         |
|---------------------------|--------------------------------------------------------------------------------------------------------------------------------------|------------------------|
| nameOverride              | Overrides name template                                                                                                              | `""`                   |
| fullnameOverride          | Overrides fullname template                                                                                                          | `""`                   |
| image.registry            | Docker registry used to pull Percona XtraDB addon image                                                                              | `stashed`              |
| image.repository          | Docker image used to backup/restore Percona XtraDB database                                                                          | `stash-percona-xtradb` |
| image.tag                 | Tag of the image that is used to backup/restore Percona XtraDB database. This is usually same as the database version it can backup. | `"5.7"`                |
| backup.args               | Arguments to pass to `mysqldump` command  during bakcup process                                                                      | `"--all-databases"`    |
| backup.socatRetry         | Optional argument sent to backup script                                                                                              | `30`                   |
| restore.args              | Arguments to pass to `mysql` command during restore process                                                                          | `""`                   |
| restore.targetAppReplicas | Optional argument sent to recovery script                                                                                            | `1`                    |
| waitTimeout               | Number of seconds to wait for the database to be ready before backup/restore process.                                                | `300`                  |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example:

```console
$ helm install stash-percona-xtradb-5.7-beta.20200708 appscode/stash-percona-xtradb -n kube-system --version=5.7-beta.20200708 --set image.registry=stashed
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```console
$ helm install stash-percona-xtradb-5.7-beta.20200708 appscode/stash-percona-xtradb -n kube-system --version=5.7-beta.20200708 --values values.yaml
```
