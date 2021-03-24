package jerakia

import (
	"encoding/json"
	"testing"

	th "github.com/jerakia/go-jerakia/testhelper"
	fake "github.com/jerakia/go-jerakia/testhelper/client"

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
