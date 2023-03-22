package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tf5server"
	"github.com/hashicorp/terraform-plugin-mux/tf5muxserver"
	"github.com/hashicorp/terraform-provider-broke/internal/newthings"
	"github.com/hashicorp/terraform-provider-broke/internal/things"
)

// Run "go generate" to format example terraform files and generate the docs for the registry/website

// If you do not have terraform installed, you can remove the formatting command, but its suggested to
// ensure the documentation is formatted properly.
//go:generate terraform fmt -recursive ./examples/

// Run the docs generation tool, check its repository for more information on how it works and how docs
// can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary
	version string = "dev"

	// goreleaser can also pass the specific commit if you want
	// commit  string = ""
)

func main() {
	var debugMode bool

	ctx := context.Background()
	flag.BoolVar(&debugMode, "debug", false, "set to true to run the things with support for debuggers like delve")
	flag.Parse()

	mainProvider := things.New(version).GRPCProvider
	otherProvider := newthings.New(version)

	muxserver, err := tf5muxserver.NewMuxServer(
		ctx,
		mainProvider,
		providerserver.NewProtocol5(otherProvider),
	)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	var opts []tf5server.ServeOpt
	if debugMode {
		opts = append(opts, tf5server.WithManagedDebug())
	}

	if err := tf5server.Serve("registry.terraform.io/notchairmk/broke", muxserver.ProviderServer, opts...); err != nil {
		log.Println(err.Error())
	}
}
