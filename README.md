# terraform-provider-jerakia

Jerakia provider for Terraform. This is a fork of the original
`jerakia/terraform-provider-jerakia`, updated for newer versions of Terraform.

## Prerequisites

* [Terraform][1]

## Terraform Configuration Example

```hcl
terraform {
  required_providers {
    jerakia = {
      source = "magicmemories/jerakia"
    }
  }
}

provider "jerakia" {
  api_url   = "https://jerakia.example.com"
  api_token = "token:token"
}

data "jerakia_lookup" "cities" {
  key       = "cities"
  namespace = "default"
}

resource "example_resource" "cities" {
  for_each = jsondecode(data.jerakia_lookup.cities.result_json)

  name       = each.key
  country    = each.value.country
  population = each.value.population
}
```

## Installation

### Using the Terraform registry (Recommended)

Simply list this provider in your module's `required_providers` block, like so:

```terraform
terraform {
  required_providers {
    jerakia = {
      source = "magicmemories/jerakia"
    }
  }
}
```

When you run `terraform init`, the provider will be downloaded automatically.

### Building from Source

> Note: The Terraform Plugin SDK supports only Go 1.15 or later (though earlier
> versions may be able to successfully compile).

1. Follow these [instructions][4] to setup a Golang development environment.
2. Check out the contents of this repository.
3. `cd` into the checked out repository
4. (Optional) Edit `Makefile` to set the version number you want assigned to
   your built binary.
5. Run `make install`

The `terraform-provider-jerakia` binary will be compiled and copied to your
[implied local mirror directory][7] for Terraform plugins. If you list it in
your `required_providers` block as described above, it will be picked up
automatically.

## Development

This project is using [Go Modules][5] for vendor support.

## Documentation

Full documentation can be found in the [`docs`][6] directory.

[1]: http://terraform.io
[4]: https://golang.org/doc/install
[5]: https://github.com/golang/go/wiki/Modules
[6]: /docs
[7]: https://www.terraform.io/docs/cli/config/config-file.html#implied-local-mirror-directories
