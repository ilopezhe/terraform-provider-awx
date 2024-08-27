//go:build mage
// +build mage

package main

import (
	"context"
	"log"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"

	// mage:import
	"github.com/ilopezhe/terraform-provider-awx/tools/kind"
	_ "github.com/ilopezhe/terraform-provider-awx/tools/kind"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func InstallAWX() error {
	log.Printf("InstallAWX to Kind Cluster")

	return sh.Run(
		"sh", "../installAwx.sh")
}

// ReCreate a kind Cluster with Awx support.
func ReCreate(ctx context.Context) {
	log.Printf("Create Kind Cluster with AWX")
	mg.CtxDeps(ctx, kind.Kind.Recreate)
	mg.CtxDeps(ctx, InstallAWX)
}

// Delete the cluster.
func Delete(ctx context.Context) {
	log.Printf("Delete Kind Cluster")
	mg.CtxDeps(ctx, kind.Kind.Delete)
}
