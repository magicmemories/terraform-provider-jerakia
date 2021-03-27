# jerakia_lookup (Data Source)

A data source for performing Jerakia lookups.

## Basic Example

```hcl
data "jerakia_lookup" "lookup_1" {
  key       = "cities"
  namespace = "example"

	metadata = {
		env    = "test"
		region = "us-west-2"
	}
}
```

## Schema

### Required

- **namespace** (String) The namespace to query.

### Optional

- **key** (String) The name of the Jerakia key to look up.
- **id** (String) The ID of this resource.
- **lookup_type** (String) The type of lookup to perform. Valid values are `first` and `cascade`.
- **merge** (String) The merge strategy to use for cascade lookups. Valid values are `array`, `hash`, and `deep_hash`.
- **metadata** (Map of String) A set of key/value pairs used by the default scope handler to set the query context.
- **policy** (String) The policy to use for the query
- **scope** (String) The name of an alternative scope handler to use.
- **scope_options** (Map of String) Options to pass to the scope handler.

### Read-Only

- **found** (Boolean) A boolean which indicates whether a result was found.
- **result_json** (String) The data returned from the query, as a JSON-encoded string.
- **status** (String) The status of the query. An error will be returned if this is `failed`.
