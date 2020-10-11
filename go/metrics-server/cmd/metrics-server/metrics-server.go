package server

import (
	"os"
	"runtime"

	"k8s.io/component-base/logs"
	genericapiserver "k8s.io/apiserver/pkg/server"
)

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	cmd := app.NewMetricsServerCommand(genericapiserver.SetupSignalHandler())
	cmd.Flags().
}
