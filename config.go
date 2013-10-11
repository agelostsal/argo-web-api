package main

import (
	"code.google.com/p/gcfg"
	"flag"
)

var flConfig = flag.String("conf", "", "specify configuration file")
var flServerPort = flag.Int("port", 0, "specify the port to listen on")
var flServerMaxProcs = flag.Int("maxprocs", 0, "specify the GOMAXPROCS")
var flMongoHost = flag.String("mongo-host", "", "specify the IP address of the MongoDB instance")
var flMongoPort = flag.Int("mongo-port", 0, "specify the port on which the MongoDB instance listens on")
var flMongoDatabase = flag.String("mongo-db", "", "specify the MongoDB database to connect to")

type Config struct {
	Server struct {
		Port     int
		Maxprocs int
		Lrucache int
	}
	MongoDB struct {
		Host string
		Port int
		Db   string
	}
}

const defaultConfig = `
    [server]
    port = 8080
    maxprocs = 4
    lrucache = 700000000

    [mongodb]
    host = "127.0.0.1"
    port = 27017
    db = "AR"
`

func LoadConfiguration() Config {
	flag.Parse()
	var cfg Config
	if *flConfig != "" {
		_ = gcfg.ReadFileInto(&cfg, *flConfig)
	} else {
		_ = gcfg.ReadStringInto(&cfg, defaultConfig)
	}

	if *flServerPort != 0 {
		cfg.Server.Port = *flServerPort
	}
	if *flServerMaxProcs != 0 {
		cfg.Server.Maxprocs = *flServerMaxProcs
	}
	if *flMongoHost != "" {
		cfg.MongoDB.Host = *flMongoHost
	}
	if *flMongoPort != 0 {
		cfg.MongoDB.Port = *flMongoPort
	}
	if *flMongoDatabase != "" {
		cfg.MongoDB.Db = *flMongoDatabase
	}

	return cfg
}
