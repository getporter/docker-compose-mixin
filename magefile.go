// +build mage

package main

import (
	"get.porter.sh/porter/mage/mixins"

	// Import common targets that all mixins should expose to the user
	// mage:import
	_ "get.porter.sh/porter/mage"
)

const (
	mixinName    = "docker-compose"
	mixinPackage = "get.porter.sh/mixin/" + mixinName
	mixinBin     = "bin/mixins/" + mixinName
)

var magefile = mixins.NewMagefile(mixinPackage, mixinName, mixinBin)

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
