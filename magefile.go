// +build mage

// This is a magefile, and is a "makefile for go".
// See https://magefile.org/
package main

import (
	"github.com/carolynvs/magex/pkg"
	"github.com/carolynvs/magex/shx"
)

var must = shx.CommandBuilder{StopOnError: true}

// Ensure mage is installed and on the PATH
func EnsureMage() error {
	return pkg.EnsureMage("")
}

// Build the docker image
func BuildImage() error {
	return shx.RunV("docker", "build", "-t=carolynvs/tabby-cat-demo-app:v0.1.0", "-f=app/Dockerfile", "app")
}

func Bundle() error {
	return shx.RunV("porter", "build", "--debug")
}
