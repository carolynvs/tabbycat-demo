// +build mage

// This is a magefile, and is a "makefile for go".
// See https://magefile.org/
package main

import (
	"github.com/magefile/mage/mg"
	"github.com/carolynvs/magex/pkg"
	"github.com/carolynvs/magex/shx"
)

const (
	version = "0.2.1"
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
	mg.SerialDeps(PublishImage, PublishBundle)
}

func PublishImage() {
	must.RunV("docker", "push", img)
}

func PublishBundle() {
	must.RunV("porter", "publish")
}
