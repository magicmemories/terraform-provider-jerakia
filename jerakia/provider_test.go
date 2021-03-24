package jerakia

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	JERAKIA_URL   = os.Getenv("JERAKIA_URL")
	JERAKIA_TOKEN = os.Getenv("JERAKIA_TOKEN")
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"jerakia": testAccProvider,
	}
}

func testAccPreCheckRequiredEnvVars(t *testing.T) {
	if JERAKIA_URL == "" {
		t.Fatal("JERAKIA_URL must be set for acceptance tests")
	}

	if JERAKIA_TOKEN == "" {
		t.Fatal("JERAKIA_TOKEN must be set for acceptance tests")
	}
}

func testAccPreCheck(t *testing.T) {
	testAccPreCheckRequiredEnvVars(t)
}
