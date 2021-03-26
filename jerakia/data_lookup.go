package jerakia

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"log"
	"strconv"

	//"github.com/jerakia/go-jerakia"
	"github.com/magicmemories/go-jerakia"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceLookup() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for performing Jerakia lookups.",

		ReadContext: dataSourceLookupRead,

		Schema: map[string]*schema.Schema{
			"key": &schema.Schema{
				Description: "The name of the Jerakia key to look up.",
				Type:         schema.TypeString,
				Required:     true,
			},

			"namespace": &schema.Schema{
				Description: "The namespace to query.",
				Type:        schema.TypeString,
				Required:    true,
			},

			"policy": &schema.Schema{
				Description: "The policy to use for the query",
				Type:        schema.TypeString,
				Optional:    true,
			},

			"lookup_type": &schema.Schema{
				Description:  "The type of lookup to perform. Valid values are `first` and `cascade`.",
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{
					"first", "cascade",
				}, false),
			},

			"merge": &schema.Schema{
				Description: "The merge strategy to use for cascade lookups. Valid values are `array`, `hash`, and `deep_hash`.",
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"array", "deep_hash", "hash",
				}, false),
			},

			"scope": &schema.Schema{
				Description: "The name of an alternative scope handler to use.",
				Type:     schema.TypeString,
				Optional: true,
			},

			"scope_options": &schema.Schema{
				Description: "Options to pass to the scope handler.",
				Type:     schema.TypeMap,
				Optional: true,
			},

			"metadata": &schema.Schema{
				Description: "A set of key/value pairs used by the default scope handler to set the query context.",
				Type:     schema.TypeMap,
				Optional: true,
			},

			// Computed
			"status": &schema.Schema{
				Description: "The status of the query. An error will be returned if this is `failed`.",
				Type:     schema.TypeString,
				Computed: true,
			},

			"found": &schema.Schema{
				Description: "A boolean which indicates whether a result was found.",
				Type:     schema.TypeBool,
				Computed: true,
			},

			"result_json": &schema.Schema{
				Description: "The data returned from the query, as a JSON-encoded string.",
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceLookupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(jerakia.Client)

	var diags diag.Diagnostics

	var lookupOpts jerakia.LookupOpts

	if v, ok := d.GetOkExists("namespace"); ok {
		namespace := v.(string)
		lookupOpts.Namespace = namespace
	}

	if v, ok := d.GetOkExists("policy"); ok {
		policy := v.(string)
		lookupOpts.Policy = policy
	}

	if v, ok := d.GetOkExists("lookup_type"); ok {
		lookupType := v.(string)
		lookupOpts.LookupType = lookupType
	}

	if v, ok := d.GetOkExists("merge"); ok {
		merge := v.(string)
		lookupOpts.Merge = merge
	}

	if v, ok := d.GetOkExists("scope"); ok {
		scope := v.(string)
		lookupOpts.Scope = scope
	}

	if v, ok := d.GetOkExists("scope_options"); ok {
		scopeOptions := expandMap(v.(map[string]interface{}))
		lookupOpts.ScopeOptions = scopeOptions
	}

	if v, ok := d.GetOkExists("metadata"); ok {
		metadata := expandMap(v.(map[string]interface{}))
		lookupOpts.Metadata = metadata
	}

	key := d.Get("key").(string)
	log.Printf("[DEBUG] jerakia_lookup lookup options for %s: %#v", key, lookupOpts)

	result, err := jerakia.Lookup(&client, key, &lookupOpts)
	if err != nil {
		return diag.FromErr(fmt.Errorf("Error querying Jerakia for %s: %s", key, err))
	}

	log.Printf("[DEBUG] jerakia_lookup result for %s: %#v", key, result)

	if result.Status == "failed" {
		return diag.FromErr(fmt.Errorf("Error querying Jerakia for %s: %s", key, result.Message))
	}

	d.SetId(generateId(lookupOpts))

	d.Set("status", result.Status)
	d.Set("found", result.Found)

	payload, err := json.Marshal(result.Payload)
	if err != nil {
		return diag.FromErr(fmt.Errorf("Error marshaling Jerakia result for %s: %s", key, err))
	}

	d.Set("result_json", string(payload))

	return diags
}

func generateId(opts jerakia.LookupOpts) string {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(opts)
	return strconv.Itoa(hash(buf.String()))
}

func expandMap(v map[string]interface{}) map[string]string {
	vmap := make(map[string]string)

	for k, v := range v {
		if value, ok := v.(string); ok && value != "" {
			vmap[k] = value
		}
	}

	return vmap
}

func hash(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	// v == MinInt
	return 0
}
