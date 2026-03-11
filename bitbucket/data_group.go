package bitbucket

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataGroup() *schema.Resource {
	return &schema.Resource{
		ReadWithoutTimeout: dataReadGroup,

		DeprecationMessage: "The Bitbucket 1.0 Groups API has been permanently deprecated by Atlassian " +
			"with no replacement. This data source will be removed in a future major version of the provider. " +
			"Remove it from your configuration.",

		Schema: map[string]*schema.Schema{
			"workspace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"slug": {
				Type:     schema.TypeString,
				Required: true,
			},
			"auto_add": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"permission": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"email_forwarding_disabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataReadGroup(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[WARN] data.bitbucket_group: Bitbucket 1.0 Groups API is permanently deprecated, removing from state")
	d.SetId("")
	return nil
}
