//go:build mage
// +build mage

package main

import (
	"get.porter.sh/magefiles/git"
	"get.porter.sh/magefiles/mixins"
)

const (
	mixinName    = "docker-compose"
	mixinPackage = "get.porter.sh/mixin/" + mixinName
	mixinBin     = "bin/mixins/" + mixinName
)

var magefile = mixins.NewMagefile(mixinPackage, mixinName, mixinBin)

func ConfigureAgent() {
	magefile.ConfigureAgent()
}

// Build the mixin
func Build() {
	magefile.Build()
}

// Cross-compile the mixin before a release
func XBuildAll() {
	magefile.XBuildAll()
}

// Run unit tests
func TestUnit() {
	magefile.TestUnit()
}

func Test() {
	magefile.Test()
}

// Publish the cross-compiled binaries.
func Publish() {
	magefile.Publish()
}

// Publish to your GitHub forks
func TestPublish(username string) {
	magefile.TestPublish(username)
}

// Install the mixin
func Install() {
	magefile.Install()
}

// Remove generated build files
func Clean() {
	magefile.Clean()
}

// SetupDCO configures your git repository to automatically sign your commits
// to comply with our DCO
func SetupDCO() error {
	return git.SetupDCO()
}
