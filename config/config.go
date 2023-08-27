// Copyright (c) 2023, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package config contains the configuration
// structs for the GoKi tool.
package config

// TODO: make all of the target fields enums

// Config is the main config struct
// that contains all of the configuration
// options for the GoKi tool
type Config struct {

	// the name of the app/library
	Name string `desc:"the name of the app/library"`

	// the description of the app/library
	Desc string `desc:"the description of the app/library"`

	// the version of the app/library
	Version string `desc:"the version of the app/library"`

	// the configuration options for the build command
	Build Build `desc:"the configuration options for the build command"`

	// the configuration options for the colorgen command
	Colorgen Colorgen `desc:"the configuration options for the colorgen command"`

	// the configuration options for the install command
	Install Install `desc:"the configuration options for the install command"`

	// the configuration options for the log command
	Log Log `desc:"the configuration options for the log command"`
}

// Build contains the configuration options
// for the build command.
type Build struct {

	// the path of the package to build
	Package string `desc:"the path of the package to build"`

	// the target platforms to build executables for
	Platform []Platform `desc:"the target platforms to build executables for"`
}

// Colorgen contains the configuration options
// for the colorgen command.
type Colorgen struct {

	// the package in which the color schemes will be used
	Package string `desc:"the package in which the color schemes will be used"`

	// the comment for the color schemes variable
	Comment string `desc:"the comment for the color schemes variable"`
}

// Install contains the configuration options
// for the install command.
type Install struct {

	// the name/path of the package to install
	Package string `desc:"the name/path of the package to install"`

	// the target platforms to install the executables on, as a list of operating systems (this should include no more than the operating system you are on, android, and ios)
	Target []string `desc:"the target platforms to install the executables on, as a list of operating systems (this should include no more than the operating system you are on, android, and ios)"`
}

// Log contains the configuration options
// for the log command.
type Log struct {

	// the target platform to view the logs for (ios or android)
	Target string `desc:"the target platform to view the logs for (ios or android)"`

	// whether to keep the previous log messages or clear them
	Keep bool `desc:"whether to keep the previous log messages or clear them"`

	// messages not generated from your app equal to or above this log level will be shown
	All string `desc:"messages not generated from your app equal to or above this log level will be shown"`
}
