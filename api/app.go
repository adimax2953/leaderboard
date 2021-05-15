package main

import (
	"api/api/server"
	"flag"
	"log"
	"os"
)

var (
	help       = flag.Bool("help", false, "Show this help")
	address    = flag.String("address", ":9090", "TCP address to listen to")
	compress   = flag.Bool("compress", false, "Whether to enable transparent response compression")
	configFile = flag.String("conf", "config.yaml", "The server configurate file")
)

//var TR int64 = 0

func main() {
	flag.Parse()
	if *help {
		flag.Usage()
		os.Exit(0)
	}

	// 寫入 Config.yaml
	var config = server.LoadConfigFromFile(*configFile)
	if config == nil {
		log.Fatalln("load config file failed")
	}
	//--------------------------------------------------------------------------------------------

	// 啟動Server
	server.Start(server.Option{
		Address:  *address,
		Compress: *compress,
		Config:   config,
	})
}
