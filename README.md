# AWX Terraform Provider

[![Go Report Card](https://goreportcard.com/badge/github.com/ilopezhe/terraform-provider-awx)](https://goreportcard.com/report/github.com/ilopezhe/terraform-provider-awx)
[![Codecov](https://img.shields.io/codecov/c/gh/ilopezhe/terraform-provider-awx)](https://app.codecov.io/gh/ilopezhe/terraform-provider-awx)
[![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/ilopezhe/terraform-provider-awx)](go.mod)
[![GitHub](https://img.shields.io/github/license/ilopezhe/terraform-provider-awx)](LICENSE)
[![Release](https://img.shields.io/github/release/ilopezhe/terraform-provider-awx.svg)](https://github.com/ilopezhe/terraform-provider-awx/releases/latest)

An autogenerated terraform provider based on the API specifications as provided by the `/api/v2/` endpoint.

## AWX Versions

Currently, built provider versions for AWX. To see which ones are active check [versions.yaml](versions.yaml)

- 24.6.1

## TODO:

- Unit tests
- Integration tests

## Download a new version of the API

You need to spin up a version of AWX you want to download the API spec from.
Older version of AWX report incorrect API spec. So manual changes may be required to fix them.

```shell
export AWX_VERSION=24.6.1
mkdir -p resources/api/$AWX_VERSION/config resource
```

s/api/$AWX_VERSION/gen-data
cat <<EOF > resources/api/$AWX_VERSION/config/default.json
{
"api_version": "$AWX_VERSION"
}
EOF
make generate-config VERSION=$AWX_VERSION
make download-api VERSION=$AWX_VERSION
make generate-config VERSION=$AWX_VERSION

````

Check the previous version of the APIs inside the `config/types` folder to see about customization.

Build the version of the current API
-------------------------------------

```shell
make generate
````

If you want to build an API for the `24.6.1` version just run

```shell
make generate VERSION=24.6.1
```

## Setup AWX for local testing in kind

A fresh AWX instance is required for automated tests, so they can ensure terraform provider is working by targetting a live AWX instance.

```sh
make dev
make port-forward
```

### Cleanup of local environment

```sh
make dev-cleanup
make no-port-forward
```
