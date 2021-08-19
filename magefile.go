// +build mage

// This is a magefile, and is a "makefile for go".
// See https://magefile.org/
package main

import (
	"github.com/carolynvs/magex/pkg"
	"github.com/carolynvs/magex/shx"
)

const (
	version = "0.1.0"
	img     = "carolynvs/tabbycat-demo-app:v" + version
)

var must = shx.CommandBuilder{StopOnError: true}

// Ensure mage is installed and on the PATH
func EnsureMage() error {
	return pkg.EnsureMage("")
}

// Build the docker image
func BuildImage() {
	must.RunV("docker", "build", "-t="+img, "-f=app/Dockerfile", "app")
}

func Bundle() {
	must.RunV("porter", "build", "--version", version, "--debug", "--verbose")
}

func Publish() {
	must.RunV("docker", "push", img)
	must.RunV("porter", "publish")
}
