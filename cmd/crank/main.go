/*
Copyright 2020 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"github.com/alecthomas/kong"
	"github.com/spf13/afero"
)

var _ = kong.Must(&cli)

var cli struct {
	Build   buildCmd   `cmd:"" help:"Build Crossplane packages."`
	Install installCmd `cmd:"" help:"Install Crossplane packages."`
	Push    pushCmd    `cmd:"" help:"Push Crossplane packages."`
}

func main() {
	buildChild := &buildChild{
		fs: afero.NewOsFs(),
	}
	pushChild := &pushChild{
		fs: afero.NewOsFs(),
	}
	ctx := kong.Parse(&cli,
		kong.Name("kubectl crossplane"),
		kong.Description("A command line tool for interacting with Crossplane."),
		// Binding a variable to kong context makes it available to all commands
		// at runtime.
		kong.Bind(buildChild, pushChild),
		kong.UsageOnError())
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
