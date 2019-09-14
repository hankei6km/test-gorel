// +build mage

package main

import (
	"fmt"
	"github.com/magefile/mage/mg"
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
	mg.Deps(Credits)
	fmt.Println("Building...")
	cmd := exec.Command("goreleaser", "--snapshot", "--skip-publish", "--rm-dist")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return err
	}
	return cmd.Wait()
}

func Credits() error {
	// mg.Deps(InstallDeps)
	gocreditsVersion := "v0.0.6"
	gocreditsFileName := fmt.Sprintf("gocredits_%s_linux_amd64", gocreditsVersion)
	gocreditsUrl := fmt.Sprintf("https://github.com/Songmu/gocredits/releases/download/%s/%s.tar.gz", gocreditsVersion, gocreditsFileName)
	creditsCmd := fmt.Sprintf(`curl -sL %s | tar --strip-components=1 --wildcards -xzf - "*/gocredits" && go mod tidy &&  ./gocredits ../my_cmd > ../CREDITS 2> /dev/null && rm gocredits`, gocreditsUrl)
	tmpDir := filepath.Join(cwd, "tmp")

	fmt.Println("Creating...")
	if _, err := os.Stat(tmpDir); err != nil {
		if err := os.Mkdir(tmpDir, os.ModePerm); err != nil {
			return err
		}
	}
	cmd := exec.Command("sh", "-c", creditsCmd)
	cmd.Dir = filepath.Join(cwd, "tmp")
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
	os.RemoveAll("tmp")
}
