/*
 * Copyright (c) 2023 Red Engine Games, LLC.
 * All Rights Reserved.
 */

package hyper_cli

import (
	"bufio"
	"bytes"
	"embed"
	"fmt"
	"github.com/briandowns/spinner"
	"io"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

//go:embed all:template
var assets embed.FS
var Version string = "dev"

type HyperEngine struct {
	Assets        map[string]string
	Configuration HyperConfiguration
}

var s *spinner.Spinner

func NewHyperEngine(configuration HyperConfiguration) *HyperEngine {
	s = spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Start()
	s.Stop()
	s.FinalMSG = "Complete"
	s.HideCursor = true
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
		_ = os.MkdirAll(folder, 0755)
		err := os.WriteFile(p, data, 0755)
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
func run_cmd_error(path string, cmd string, args ...string) string {
	exe := exec.Command(cmd, args...)
	exe.Dir = path
	out, err := exe.StdoutPipe()
	if err != nil {
		log.Printf("[ERROR] %v", err)
	}
	buf := new(bytes.Buffer)
	exe.Run()
	_, _ = io.Copy(buf, out)
	return buf.String()
}
func (h *HyperEngine) NewProject(name string) error {
	s.Suffix = " New Hyper Pipeline"
	s.Start()
	dir, _ := os.Getwd()
	if name != "." {
		_ = os.MkdirAll(name, 0755)
		dir = filepath.Join(dir, name)
	}
	run_cmd(dir, "git", "init")
	out := run_cmd_error(dir, "npm", "init", "-y")
	out = out + run_cmd_error(dir, "npm", "i", "vite@latest", "mdb-ui-kit@latest", "sass")
	WriteAssets(dir)
	fmt.Println(out)
	s.Stop()
	return nil
}
func Install(args ...string) {
	s.Suffix = " Package Installation"
	s.Start()
	dir, _ := os.Getwd()
	fmt.Println(run_cmd_error(dir, "npm", args...))
	s.Stop()
}
func (h *HyperEngine) Build() error {
	s.Suffix = " Compilation"
	s.Start()
	dir, _ := os.Getwd()
	outDir := fmt.Sprintf("--outDir=%s", filepath.Join(dir, h.Configuration.OutputDir))
	out := run_cmd_error(dir, "npx", "vite", "build", "--emptyOutDir", outDir)
	s.Stop()
	fmt.Println(out)
	return nil
}

func (h *HyperEngine) Serve() error {
	fmt.Println("Starting Service")
	e := h.Build()
	if e != nil {
		return e
	}
	cmd := exec.Command("npx", "vite", "--port", fmt.Sprintf("%d", h.Configuration.DevPort))
	stdout, _ := cmd.StderrPipe()
	e = cmd.Start()
	if e != nil {
		return e
	}
	fmt.Println("HyperNet Active: http://localhost:8080")
	fmt.Println("")
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
