---
title: Backup & Restore Clustered Percona XtraDB Database | Stash
description: Backup Clustered Percona XtraDB Database database using Stash
menu:
  product_stash_{{ .version }}:
    identifier: percona-xtradb-clustered-guide-{{ .subproject_version }}
    name: Backup & Restore Clustered Percona XtraDB
    parent: stash-percona-xtradb-guides-{{ .subproject_version }}
    weight: 20
product_name: stash
menu_name: product_stash_{{ .version }}
section_menu_id: stash-addons
---

**Deploy Database:**

```bash
$ kubectl apply -f ./docs/examples/clustered/backup/sample-xtradb-cluster.yaml
perconaxtradb.kubedb.com/sample-xtradb-cluster created
```

**Create Secret:**
```bash
$ kubectl create secret generic -n demo gcs-secret \
                --from-file=./RESTIC_PASSWORD \
                --from-file=./GOOGLE_PROJECT_ID \
                --from-file=./GOOGLE_SERVICE_ACCOUNT_JSON_KEY
```

**Create Repository:**

```bash
$ kubectl apply -f ./docs/examples/clustered/backup/repository.yaml
repositories.stash.appscode.com/gcs-repo-xtradb-cluster created
```

**Create BackupConfiguration:**

```bash
$ kubectl apply -f ./docs/examples/clustered/backup/backupconfiguration.yaml 
backupconfiguration.stash.appscode.com/sample-xtradb-cluster-backup created
```

## Restore

**Create Restored Database:**

```bash
$ kubectl apply -f ./docs/examples/clustered/restore/restored-xtradb-cluster.yaml 
perconaxtradb.kubedb.com/restored-xtradb-cluster created
```

**Create RestoreSession:**
```bash
$ kubectl apply -f ./docs/examples/clustered/restore/restoresession.yaml
restoresession.stash.appscode.com/restored-xtradb-cluster-restore created
```