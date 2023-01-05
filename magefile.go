//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/hankei6km/go-ac"
	"github.com/magefile/mage/mg"
)

var cwd = func() string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return cwd
}()

// var mainDir = filepath.Join(cwd, "my_cmd")
// var tmpDir = filepath.Join(cwd, "tmp")
var mainDir = "my_cmd"
var tmpDir = "tmp"

var workDir = filepath.Join(cwd, tmpDir)

var pkgs = []string{"my_cmd"}

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

func Snapshot() error {
	fmt.Println("Building...")
	cmd := exec.Command("goreleaser", "--snapshot", "--skip-publish", "--rm-dist", "-p", "1")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return err
	}
	if err := cmd.Wait(); err != nil {
		return err
	}

	// if err := os.Remove("CREDITS"); err != nil {
	// 	return err
	// }
	return nil
}

func TmpDir() error {
	fmt.Println("Creating tmp directory...")
	if _, err := os.Stat(tmpDir); err != nil {
		if err := os.Mkdir(tmpDir, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func Credits() error {
	mg.Deps(TmpDir)

	fmt.Println("Writing CREDITS...")
	b := ac.NewOutputBuilder().
		GoSumFile(filepath.Join(cwd, "go.sum"))
	d := ac.NewDistBuilder().
		DistDir(filepath.Join(cwd, "dist")).
		WorkDir(workDir).
		OutDir(cwd).
		OutputBuilder(b).
		Build()
	if err := d.Run(); err != nil {
		return err
	}
	return nil
}

func Test() error {
	// mg.Deps(InstallDeps)
	fmt.Println("Testing...")
	for _, p := range pkgs {
		cmd := exec.Command("go", "test", "-v", "-race", "-count", "5", "-failfast")
		cmd.Dir = filepath.Join(cwd, p)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Start(); err != nil {
			return err
		}
		if err := cmd.Wait(); err != nil {
			return err
		}
	}
	return nil
}

func cleanCredits() error {
	files, err := filepath.Glob(filepath.Join(cwd, "CREDITS_*"))
	if err != nil {
		return err
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			return err
		}
	}
	return nil
}

func Clean() {
	fmt.Println("Cleaning...")
	os.RemoveAll("dist")
	os.RemoveAll("tmp")
	os.Remove("CREDITS")
	cleanCredits()
}
