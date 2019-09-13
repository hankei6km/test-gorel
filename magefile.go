// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

var cwd = func() string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return cwd
}()

var pkgs = []string{"my_cmd"}

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

func Snapshot() error {
	// mg.Deps(InstallDeps)
	fmt.Println("Building...")
	cmd := exec.Command("goreleaser", "--snapshot", "--skip-publish", "--rm-dist")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return err
	}
	return cmd.Wait()
}

func Release() error {
	// mg.Deps(InstallDeps)
	fmt.Println("Building...")
	cmd := exec.Command("goreleaser", "--snapshot", "--rm-dist")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return err
	}
	return cmd.Wait()
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

func Clean() {
	fmt.Println("Cleaning...")
	os.RemoveAll("dist")
}
