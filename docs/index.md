# Jerakia Provider

Use Terraform to query Jerakia

## Example Usage

To configure this provider, do the following:

```hcl
provider "jerakia" {
  api_url   = "https://jerakia.example.com"
  api_token = "tokentoken"
}
```

## Schema

* `api_url` - *Required* - The URL to the Jerakia service. This can also be set
  with the `JERAKIA_URL` environment variable.

* `api_token` - *Required* - The token to authenticate with. This can also be
	set with the `JERAKIA_TOKEN` environment variable.
