package jerakia

import (
	"encoding/json"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceLookup_basic(t *testing.T) {
	expectedPayload := map[string]interface{}{"argentina": "buenos aires", "france": "paris", "spain": "malaga"}
	expectedJSON, err := json.Marshal(expectedPayload)
	if err != nil {
		t.Fatalf("Unable to marshal JSON: %s", err)
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceLookup_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.jerakia_lookup.lookup_1", "result_json", string(expectedJSON)),
				),
			},
		},
	})
}

func TestAccDataSourceLookup_singleBool(t *testing.T) {
	expectedJSON := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceLookup_singleBool,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.jerakia_lookup.lookup_1", "result_json", string(expectedJSON)),
				),
			},
		},
	})
}

func TestAccDataSourceLookup_metadata(t *testing.T) {
	expectedPayload := []interface{}{"bob", "lucy", "david"}
	expectedJSON, err := json.Marshal(expectedPayload)
	if err != nil {
		t.Fatalf("Unable to marshal JSON: %s", err)
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceLookup_metadata,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.jerakia_lookup.lookup_1", "result_json", string(expectedJSON)),
				),
			},
		},
	})
}

func TestAccDataSourceLookup_hash(t *testing.T) {
	expectedPayload := map[string]interface{}{
		"key0": map[string]interface{}{
			"element0": "common",
		},
		"key1": map[string]interface{}{
			"element2": "env",
		},
		"key2": map[string]interface{}{
			"element3": map[string]interface{}{
				"subelement3": "env",
			},
		},
		"key3": "env",
	}
	expectedJSON, err := json.Marshal(expectedPayload)
	if err != nil {
		t.Fatalf("Unable to marshal JSON: %s", err)
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceLookup_hash,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.jerakia_lookup.lookup_1", "result_json", string(expectedJSON)),
				),
			},
		},
	})
}

func TestAccDataSourceLookup_keyless(t *testing.T) {
	expectedPayload := map[string]interface{}{"foo": "bar", "hello": "world"}
	expectedJSON, err := json.Marshal(expectedPayload)
	if err != nil {
		t.Fatalf("Unable to marshal JSON: %s", err)
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceLookup_keyless,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.jerakia_lookup.lookup_1", "result_json", string(expectedJSON)),
				),
			},
		},
	})
}


const testAccDataSourceLookup_basic = `
  data "jerakia_lookup" "lookup_1" {
    key       = "cities"
    namespace = "test"
  }
`

const testAccDataSourceLookup_singleBool = `
  data "jerakia_lookup" "lookup_1" {
    key       = "booltrue"
    namespace = "test"
  }
`

const testAccDataSourceLookup_metadata = `
  data "jerakia_lookup" "lookup_1" {
    key       = "users"
    namespace = "test"

    metadata = {
      hostname = "example"
    }
  }
`

const testAccDataSourceLookup_hash = `
  data "jerakia_lookup" "lookup_1" {
    key       = "hash"
    namespace = "test"

		lookup_type = "cascade"
		merge       = "hash"

    metadata = {
      env = "dev"
    }
  }
`

const testAccDataSourceLookup_keyless = `
  data "jerakia_lookup" "lookup_1" {
    namespace = "keyless"
  }
`