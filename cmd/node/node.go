// Command gocron-node
package main

import (
	"flag"
	"github.com/palemoonshine1224/czjCron/internal/modules/rpc/server"
	log "github.com/sirupsen/logrus"
	"os"
	"runtime"
)

func main() {
	var serverAddr string
	var allowRoot bool
	var version bool
	var logLevel string
	flag.BoolVar(&allowRoot, "allow-root", false, "./gocron-node -allow-root")
	flag.StringVar(&serverAddr, "s", "0.0.0.0:5922", "./gocron-node -s ip:port")
	flag.BoolVar(&version, "v", false, "./gocron-node -v")
	flag.StringVar(&logLevel, "log-level", "info", "-log-level error")
	flag.Parse()
	level, err := log.ParseLevel(logLevel)
	if err != nil {
		log.Fatal(err)
	}
	log.SetLevel(level)

	if runtime.GOOS != "windows" && os.Getuid() == 0 && !allowRoot {
		log.Fatal("Do not run gocron-node as root user")
		return
	}

	server.Start(serverAddr)
}
