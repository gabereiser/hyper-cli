/*
 * Copyright (c) 2023 Red Engine Games, LLC.
 * All Rights Reserved.
 */

package main

import (
	"flag"
	"fmt"
	hyper_cli "hyper-cli"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"
)

var banner = fmt.Sprintf(`
██╗  ██╗██╗   ██╗██████╗ ███████╗██████╗ 
██║  ██║╚██╗ ██╔╝██╔══██╗██╔════╝██╔══██╗
███████║ ╚████╔╝ ██████╔╝█████╗  ██████╔╝
██╔══██║  ╚██╔╝  ██╔═══╝ ██╔══╝  ██╔══██╗
██║  ██║   ██║   ██║     ███████╗██║  ██║
╚═╝  ╚═╝   ╚═╝   ╚═╝     ╚══════╝╚═╝  ╚═╝
(c) 2020-%d HyperCloud.network`, time.Now().Year())

func main() {
	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		os.Exit(1)
	} else {
		switch strings.TrimSpace(strings.ToLower(args[0])) {
		case "build":
			h := hyper_cli.NewHyperEngine(hyper_cli.LoadConfiguration())
			err := h.Build()
			if err != nil {
				panic(fmt.Errorf("[ERROR] %v", err))
			}
			log.Println("[SUCCESS]")
			os.Exit(0)
		case "serve":
			end := make(chan os.Signal, 1)
			signal.Notify(end, os.Interrupt)
			h := hyper_cli.NewHyperEngine(hyper_cli.LoadConfiguration())
			go func() {
				<-end
				fmt.Println("Interrupt detected")
				os.Exit(0)
			}()
			err := h.Serve()
			if err != nil {
				panic(fmt.Errorf("[ERROR] %v", err))
			}
			os.Exit(0)
		case "init":
			if len(args) != 2 {
				panic(fmt.Errorf("[ERROR] You must supply a project name or '.' see help for details"))
			}
			h := hyper_cli.NewHyperEngine(hyper_cli.DefaultConfiguration())
			err := h.NewProject(args[1])
			if err != nil {
				panic(fmt.Errorf("[ERROR] %v", err))
			}
			os.Exit(0)
		case "info":
			hyper_cli.PrintAssets()
			os.Exit(0)
		case "install":
			hyper_cli.Install(args...)
			os.Exit(0)
		default:
			flag.Usage()
			os.Exit(0)
		}
	}
}

func init() {
	fmt.Println(banner)
	fmt.Println()
	fmt.Println(hyper_cli.Version)
	flag.Usage = func() {
		fmt.Println("\r\nUsage: hyper [command]")
		fmt.Println("-----------------------------")
		fmt.Println("\tinit [dir]\tInits a new project")
		fmt.Println("\tbuild\t\tBuilds the current project.")
		fmt.Println("\tserve\t\tServes the current project for development.")
		fmt.Println("\thelp\t\tthis help message.")
		flag.PrintDefaults()
	}
	flag.Parse()
}
