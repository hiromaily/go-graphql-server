package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"

	"github.com/hiromaily/go-graphql-server/pkg/config"
	"github.com/hiromaily/go-graphql-server/pkg/files"
)

var (
	tomlPath  = flag.String("toml", "", "TOML file path")
	isVersion = flag.Bool("v", false, "version")
	// -d daemon mode
	version string
)

var usage = `Usage: %s [options...]
Options:
  -toml      Toml file path for config
  -v         show version
`

// init() can not be used because it affects main_test.go as well.
func init() {
}

func parseFlag() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage, os.Args[0])
	}
	flag.Parse()
}

func checkVersion() {
	if *isVersion {
		fmt.Printf("%s %s\n", "book-teacher", version)
		os.Exit(0)
	}
}

func getConfig() *config.Root {
	configPath := files.GetConfigPath(*tomlPath)
	if configPath == "" {
		log.Fatal(errors.New("config file is not found"))
	}
	log.Println("config file: ", configPath)
	conf, err := config.NewConfig(configPath)
	if err != nil {
		panic(err)
	}
	return conf
}

func main() {
	parseFlag()
	checkVersion()

	conf := getConfig()
	regi := NewRegistry(conf)

	srv := regi.NewServer()
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
	srv.Clean()
}
