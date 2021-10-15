package main

import (
	"fmt"
	"path"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

const name string = "ambilight"
const buildDir string = "build"

type (
	Build   mg.Namespace
	Test    mg.Namespace
	Docker  mg.Namespace
	CI      mg.Namespace
	Install mg.Namespace
)

type BuildOptions struct {
	Static    bool
	Extension string
}

func build(os string, arch string, opts BuildOptions) error {
	env := map[string]string{
		"GOOS":   os,
		"GOARCH": arch,
	}

	// Build static binary
	if opts.Static {
		env["CGO_ENABLED"] = "0"
	}

	outputName := fmt.Sprintf(
		"%s-%s-%s", name, os, arch,
	)

	outputPath := path.Join(
		buildDir,
		outputName,
	)

	if opts.Extension != "" {
		outputPath += fmt.Sprintf(".%s", opts.Extension)
	}

	args := []string{
		"build",
		"-o",
		outputPath,
		"-ldflags",
		"-s -w",
		".",
	}

	return sh.RunWith(
		env, "go", args...,
	)
}

// Builds all binaries
func (Build) All() {
	mg.Deps(
		Build.DarwinAmd64,
		Build.LinuxAmd64,
		Build.WindowsAmd64,
	)
}

// Builds the binary for Linux (amd64)
func (Build) LinuxAmd64() error {
	return build("linux", "amd64", BuildOptions{})
}

// Builds the binary for Windows (amd64)
func (Build) WindowsAmd64() error {
	return build("windows", "amd64", BuildOptions{
		Extension: "exe",
	})
}

// Builds the binary for Darwin/macOS (amd64)
func (Build) DarwinAmd64() error {
	return build("darwin", "amd64", BuildOptions{})
}


// Compresses all binaries into a single tarball
func (CI) CompressBinaries() error {
	return sh.Run("tar", "-cvzf", "binaries.tar.gz", "build")
}

// Cleans up build directories
func Clean() {
	sh.Rm("build")
}