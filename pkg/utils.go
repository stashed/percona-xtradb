package pkg

import (
	"fmt"
	"os/exec"
	"time"

	"stash.appscode.dev/stash/pkg/restic"

	"github.com/appscode/go/log"
	"k8s.io/client-go/kubernetes"
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
	catalogClient appcatalog_cs.Interface

	namespace         string
	appBindingName    string
	outputDir         string
	dumpArgs          string
	mysqlArgs         string
	garbdCnf          kubedbconfig_api.GaleraArbitratorConfiguration
	socatRetry        int32
	targetAppReplicas int32

	setupOptions  restic.SetupOptions
	backupOptions restic.BackupOptions
	dumpOptions   restic.DumpOptions
}

func waitForDBReady(host string, port int32) {
	log.Infoln("Checking database connection")
	cmd := fmt.Sprintf(`nc "%s" "%d" -w 30`, host, port)
	for {
		if err := exec.Command(cmd).Run(); err != nil {
			break
		}
		log.Infoln("Waiting... database is not ready yet")
		time.Sleep(5 * time.Second)
	}
}
