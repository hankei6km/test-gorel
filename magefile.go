// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

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

func TmpDir() error {
	fmt.Println("Creating tmp directory...")
	if _, err := os.Stat(tmpDir); err != nil {
		if err := os.Mkdir(tmpDir, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func StripTest() error {
	mg.Deps(TmpDir)
	fmt.Println("Copying main src to tmp without _test.go...")

	mainOnlyDir := filepath.Join(tmpDir, mainDir)
	if _, err := os.Stat(mainOnlyDir); err == nil {
		if err := os.RemoveAll(mainOnlyDir); err != nil {
			return err
		}
	}

	stripCmd := fmt.Sprintf(`tar --exclude="*_test.go" -cf - %s/ | tar --directory=%s -xf -`, mainDir, tmpDir)
	cmd := exec.Command("sh", "-c", stripCmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return err
	}
	if err := cmd.Wait(); err != nil {
		return err
	}
	cmd = exec.Command("sh", "-c", "go mod init && go mod tidy")
	cmd.Dir = mainOnlyDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return err
	}
	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}

func Credits() error {
	mg.Deps(StripTest)
	gocreditsVersion := "v0.0.6"
	gocreditsFileName := fmt.Sprintf("gocredits_%s_linux_amd64", gocreditsVersion)
	gocreditsUrl := fmt.Sprintf("https://github.com/Songmu/gocredits/releases/download/%s/%s.tar.gz", gocreditsVersion, gocreditsFileName)
	creditsCmd := fmt.Sprintf(`curl -sL %s | tar --strip-components=1 --wildcards -xzf - "*/gocredits" &&  ./gocredits ./%s > ../CREDITS 2> /dev/null && rm gocredits`, gocreditsUrl, mainDir)

	fmt.Println("Creating...")
	cmd := exec.Command("sh", "-c", creditsCmd)
	cmd.Dir = tmpDir
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
