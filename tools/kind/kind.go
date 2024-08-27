// Package kind commands
package kind

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

const waitTime = 30

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Kind Mage Command Namespace.
type Kind mg.Namespace

// Delete the cluster.
func (Kind) Delete() error {
	log.Printf("Delete Kind Cluster")

	return sh.Run(
		"kind", "delete", "cluster")
}

// Recreate The full Cluster.
func (Kind) Recreate(ctx context.Context) {
	log.Printf("Recreate Kind Cluster")
	mg.CtxDeps(ctx, Kind.Delete)
	mg.CtxDeps(ctx, Kind.Create)
}

// Create a kind Cluster with Ingress support.
func (Kind) Create(ctx context.Context) {
	log.Printf("Create Kind Cluster with Ingress")
	mg.CtxDeps(ctx, Kind.InstallKind)
	mg.CtxDeps(ctx, Kind.InstallIngress)
}

// InstallIngress to Cluster.
func (Kind) InstallIngress() error {
	log.Printf("Install Ingress to Cluster")

	nginxVersion := "main"

	urlReleaseBase := fmt.Sprintf(
		"https://raw.githubusercontent.com/kubernetes/ingress-nginx/%s",
		nginxVersion,
	)

	url := fmt.Sprintf("%s/deploy/static/provider/kind/deploy.yaml", urlReleaseBase)
	err := sh.Run("kubectl", "apply", "-f", url, "--wait=true")
	check(err)
	time.Sleep(waitTime * time.Second)

	return sh.Run("kubectl", "wait", "--namespace", "ingress-nginx",
		"--for=condition=ready", "pod",
		"--selector=app.kubernetes.io/component=controller",
		"--timeout=240s")
}

// InstallKind to local System.
func (Kind) InstallKind() error {
	log.Printf("Create Cluster")

	kindConfig := `
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
  kubeadmConfigPatches:
  - |
    kind: InitConfiguration
    nodeRegistration:
      kubeletExtraArgs:
        node-labels: "ingress-ready=true"
  extraPortMappings:
  - containerPort: 80
    hostPort: 80
    protocol: TCP
  - containerPort: 443
    hostPort: 443
    protocol: TCP
`

	d1 := []byte(kindConfig)
	configPath := path.Join(os.TempDir(), "kindconfig.yaml")
	err := os.WriteFile(configPath, d1, 0o600)

	defer os.Remove(configPath)
	check(err)

	// check kind allways exists
	out, err := sh.Output("kind", "get", "clusters", "-q")
	check(err)

	if strings.Contains(out, "kind") {
		return sh.Run("kind", "export", "kubeconfig")
	}

	return sh.Run("kind", "create", "cluster", fmt.Sprintf("--config=%s", configPath))
}
