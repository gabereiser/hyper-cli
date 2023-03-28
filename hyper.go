/*
 * Copyright (c) 2023 Red Engine Games, LLC.
 * All Rights Reserved.
 */

package hyper_cli

import (
	"bufio"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//go:embed all:template
var assets embed.FS
var Version string = "dev"

type HyperEngine struct {
	Assets        map[string]string
	Configuration HyperConfiguration
}

func NewHyperEngine(configuration HyperConfiguration) *HyperEngine {
	return &HyperEngine{
		Assets:        make(map[string]string),
		Configuration: configuration,
	}
}
func getAllFilenames(efs *embed.FS) (files []string, err error) {
	if err := fs.WalkDir(efs, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		files = append(files, path)

		return nil
	}); err != nil {
		return nil, err
	}

	return files, nil
}
func PrintAssets() {
	files, err := getAllFilenames(&assets)
	if err != nil {
		panic(fmt.Errorf("[ERROR] %v", err))
	}
	for _, file := range files {
		fmt.Println(file)
	}
}
func WriteAssets(path string) {
	files, err := getAllFilenames(&assets)
	if err != nil {
		panic(fmt.Errorf("[ERROR] %v", err))
	}
	for _, file := range files {
		data, _ := assets.ReadFile(file)
		file = strings.TrimPrefix(file, "template/")
		p := filepath.Join(path, file)
		folder := filepath.Dir(p)
		_ = os.MkdirAll(folder, 0644)
		err := os.WriteFile(p, data, 0644)
		if err != nil {
			panic(fmt.Errorf("[ERROR] %v", err))
		}
	}
}
func run_cmd(path string, cmd string, args ...string) string {
	exe := exec.Command(cmd, args...)
	exe.Dir = path
	out, err := exe.CombinedOutput()
	if err != nil {
		log.Printf("[ERROR] %v", err)
	}
	return string(out)
}
func (h *HyperEngine) NewProject(name string) error {
	dir, _ := os.Getwd()
	if name != "." {
		_ = os.MkdirAll(name, 0644)
		dir = filepath.Join(dir, name)
	}
	run_cmd(dir, "git", "init")
	out := run_cmd(dir, "npm", "init", "-y")
	out = out + run_cmd(dir, "npm", "i", "vite@latest", "mdb-ui-kit@latest", "sass")
	WriteAssets(dir)
	fmt.Println(out)
	return nil
}
func Install(args ...string) {
	dir, _ := os.Getwd()
	fmt.Println(run_cmd(dir, "npm", args...))
}
func (h *HyperEngine) Build() error {
	dir, _ := os.Getwd()
	outDir := fmt.Sprintf("--outDir=%s", filepath.Join(dir, h.Configuration.OutputDir))
	out := run_cmd(dir, "npx", "vite", "build", "--emptyOutDir", outDir)
	fmt.Println(out)
	return nil
}

func (h *HyperEngine) Serve() error {
	e := h.Build()
	if e != nil {
		return e
	}
	cmd := exec.Command("npx", "vite", "preview", "--port", fmt.Sprintf("%d", h.Configuration.DevPort))
	stdout, _ := cmd.StdoutPipe()
	e = cmd.Start()
	if e != nil {
		return e
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}
	e = cmd.Wait()
	if e != nil {
		return e
	}
	return nil
}
