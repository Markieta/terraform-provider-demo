# terraform-provider-demo

Demonstrating a custom Terraform provider that creates/updates file content inside a GitHub repository.

## Disclaimer

**This custom provider is not recommended for production environments.**

I have written this provider for demonstration purposes, there is still room for improvement.

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

As opposed to exposing your GitHup Personal Access Token inside your Terraform modules, you can export it as an environment variable:

```bash
export GITHUB_TOKEN="<GITHUB_TOKEN>"
```

### Terraform Usage

This directory contains a simple `main.tf` that defines the custom resource, `demo_repo_content`, to be created on `apply`. Modify the `main.tf` to configure your GitHub repository details.

Initialize the custom provider (and any other providers) to be used here, then run `apply`:

```bash
terraform init
terraform apply
```

You will then be prompted to review and confirm that `demo_repo_content.main` will be created.

If everything succeeds, the file with the content defined in your Terraform module(s) will be created/updated in your desired GitHub repository and branch.

Refer to [Installing Plugins](https://www.terraform.io/docs/plugins/basics.html#installing-plugins) for configuring this custom provider.
