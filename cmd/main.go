package main

import (
	"errors"
	"flag"
	"log"
	"os"
	"path"
	"shared-place/src/types"
)

func main() {
	// Flags
	var cfgDir string
	flag.StringVar(&cfgDir, "cfg", "", "Set config file directory")
	flag.Parse()

	if cfgDir == "" {
		log.Fatal("Invalid argument! Use --help to list flags.")
	}

	// Check cfg file
	fi, err := os.Stat(cfgDir)
	if errors.Is(err, os.ErrNotExist) {
		log.Fatal("Failed to open the configuration file.")
	} else if err != nil {
		log.Fatal("An error occured! Please report this issue on our repo issues.")
	}

	ext := path.Ext(fi.Name())
	switch ext {
	case ".yaml":
		break
	case ".yml":
		break
	case ".json":
		break
	case "toml":
		break
	default:
		log.Fatal("This file extension is not supported.")
	}

	// Connect to the server
	s, err := types.NewServer("123.456.789.10", 8080, "/var/contents")
	if err != nil {
		log.Fatal(err)
	}
	s.Connect()
}
