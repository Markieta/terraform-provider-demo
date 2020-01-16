# terraform-provider-demo

Demonstrating a custom Terraform provider.

## Environment Details

The following instructions are tested and based on:

* Fedora 31
* Terraform 0.12.19
* Go 1.13.5

## Prerequisites

Some initial requirements.

### Terraform Setup

If Terraform is not already installed:

```bash
# or find the latest: https://releases.hashicorp.com/terraform/
wget https://releases.hashicorp.com/terraform/0.12.19/terraform_0.12.19_linux_amd64.zip
unzip terraform_0.12.19_linux_amd64.zip
sudo unzip terraform_0.12.19_linux_amd64.zip -d /usr/local/bin/
```

### Go Setup

Firstly, install Go:

```bash
sudo dnf install golang
```

> [Terraform Core and Terraform provider development requires using the Go toolchain in the new "modules mode", which in current versions of Go is not the default. (Atkins, 2019)](https://stackoverflow.com/a/58701749/1154847)

Execute the following to enable the modules mode:

```bash
go mod init mymodule/test
```

The Go build will now be able to download all required dependencies.

## Usage

How to use this custom Terraform provider.

### Building Source

Whenever changes are made to the provider, rebuild the binary:

```bash
go build -o terraform-provider-demo
```

### Authentication

