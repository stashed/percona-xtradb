/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Free Trial License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Free-Trial-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package pkg

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"

	stash "stash.appscode.dev/apimachinery/client/clientset/versioned"
	"stash.appscode.dev/apimachinery/pkg/restic"

	shell "gomodules.xyz/go-sh"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
	kmapi "kmodules.xyz/client-go/api/v1"
	appcatalog "kmodules.xyz/custom-resources/apis/appcatalog/v1alpha1"
	appcatalog_cs "kmodules.xyz/custom-resources/client/clientset/versioned"
	kubedbconfig_api "kubedb.dev/apimachinery/apis/config/v1alpha1"
)

const (
	mySqlUser     = "username"
	mySqlPassword = "password"

	mySqlDumpFile        = "dumpfile.sql"
	xtraBackupStreamFile = "xtrabackup.stream"

	mySqlDumpCMD     = "mysqldump"
	mySqlRestoreCMD  = "mysql"
	envMySqlPassword = "MYSQL_PWD"
	defaultDumpArgs  = "--all-databases"
)

type perconaOptions struct {
	kubeClient    kubernetes.Interface
	stashClient   stash.Interface
	catalogClient appcatalog_cs.Interface

	namespace           string
	backupSessionName   string
	appBindingName      string
	appBindingNamespace string
	outputDir           string
	storageSecret       kmapi.ObjectReference
	xtradbArgs          string
	garbdCnf            kubedbconfig_api.GaleraArbitratorConfiguration
	socatRetry          int32
	targetAppReplicas   int32
	waitTimeout         int32

	setupOptions  restic.SetupOptions
	backupOptions restic.BackupOptions
	dumpOptions   restic.DumpOptions
}
type sessionWrapper struct {
	sh  *shell.Session
	cmd *restic.Command
}

func (opt *perconaOptions) newSessionWrapper() *sessionWrapper {
	return &sessionWrapper{
		sh:  shell.NewSession(),
		cmd: &restic.Command{},
	}
}

func (session *sessionWrapper) setDatabaseCredentials(kubeClient kubernetes.Interface, appBinding *appcatalog.AppBinding) error {
	appBindingSecret, err := kubeClient.CoreV1().Secrets(appBinding.Namespace).Get(context.TODO(), appBinding.Spec.Secret.Name, metav1.GetOptions{})
	if err != nil {
		return err
	}

	err = appBinding.TransformSecret(kubeClient, appBindingSecret.Data)
	if err != nil {
		return err
	}
	session.sh.SetEnv(envMySqlPassword, string(appBindingSecret.Data[mySqlPassword]))
	session.cmd.Args = append(session.cmd.Args, "-u", string(appBindingSecret.Data[mySqlUser]))
	return nil
}

func (session *sessionWrapper) setDatabaseConnectionParameters(appBinding *appcatalog.AppBinding) error {
	hostname, err := appBinding.Hostname()
	if err != nil {
		return err
	}
	session.cmd.Args = append(session.cmd.Args, "-h", hostname)

	port, err := appBinding.Port()
	if err != nil {
		return err
	}
	// if port is specified, append port in the arguments
	if port != 0 {
		session.cmd.Args = append(session.cmd.Args, fmt.Sprintf("--port=%d", port))
	}
	return nil
}

func (session *sessionWrapper) setUserArgs(args string) {
	for _, arg := range strings.Fields(args) {
		session.cmd.Args = append(session.cmd.Args, arg)
	}
}

func waitForDBReady(appBinding *appcatalog.AppBinding, waitTimeout int32) error {
	klog.Infoln("Checking database connection")
	hostname, err := appBinding.Hostname()
	if err != nil {
		return err
	}

	port, err := appBinding.Port()
	if err != nil {
		return err
	}
	cmd := fmt.Sprintf(`nc "%s" "%d" -w %d`, hostname, port, waitTimeout)
	for {
		if err := exec.Command(cmd).Run(); err != nil {
			break
		}
		klog.Infoln("Waiting... database is not ready yet")
		time.Sleep(5 * time.Second)
	}
	return nil
}
