# Terraform Provider Forefront

This is a custom Terraform Provider containing utility functions (like deep map/object string replacement) for advanced Terraform configurations. It is built using the modern `terraform-plugin-framework`.

## Included Functions

- `provider::forefront::lookup(key, map, default)`
  - Returns a value from a map given a key. If the key is not found, returns the `default`.
- `provider::forefront::replace_deep(collection, search, replace)`
  - Recursively walks through any complex object (maps, lists, sets, tuples, objects) to find all string values and replaces all occurrences of `search` with `replace`.

## Setup and Usage

To use this provider in your Terraform configuration, add the following to your `.tf` file:

```hcl
terraform {
  required_providers {
    forefront = {
      source = "magnus-ffcg/forefront"
    }
  }
}

provider "forefront" {}

output "example" {
  value = provider::forefront::replace_deep({
    "key1" = "hello word"
  }, "word", "world")
}
```

## Local Development Requirements

- [Go](https://golang.org/doc/install) >= 1.22
- [Terraform](https://developer.hashicorp.com/terraform/downloads) >= 1.5

### Testing Locally (Dev Overrides)

During development, you can test the provider locally without publishing it to a registry. 
Compile the binary in the root directory:

```bash
go build -o terraform-provider-forefront
```

Then, configure your `~/.terraformrc` to use the local binary:

```hcl
provider_installation {
  dev_overrides {
    "magnus-ffcg/forefront" = "/Users/<your-user>/path/to/forefront-provider"
  }
  direct {}
}
```

Now you can run `terraform plan` in the `examples/` directory to see it in action.

## Deployment / CI pipelines

This provider is configured to use [GoReleaser](https://goreleaser.com/) for building, packing, and publishing the provider binaries to GitHub Releases. The `.goreleaser.yml` at the root configures the cross-compilation matrix.

When you are ready to publish a new version of the provider:
1. Tag the repository with a semantic version: `git tag v1.0.0`
2. Push the tag: `git push origin v1.0.0`
3. Run GoReleaser (either locally with `goreleaser release` or via GitHub actions).

If you are publishing to the public HashiCorp Registry, the Registry will automatically pick up the new GitHub release.

## Generating Documentation

Run `go generate ./...` to auto-generate the markdown files in the `docs/` directory representing the provider variables and functions. These docs are directly displayed on the Terraform Registry.
