package main

import (
	"context"
	"flag"
	"log"

	"github.com/ilopezhe/terraform-provider-awx/internal/awx"
	"github.com/ilopezhe/terraform-provider-awx/version"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"github.com/ilopezhe/terraform-provider-awx/internal/provider"
)

func main() {
	var debug bool
	var err error

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	if err = providerserver.Serve(
		context.Background(),
		provider.NewFuncProvider(version.Version, nil, awx.Resources(), awx.DataSources()),
		providerserver.ServeOpts{
			Address: "registry.terraform.io/ilopezhe/awx",
			Debug:   debug,
		},
	); err != nil {
		log.Fatal(err.Error())
	}
}
