package bitbucket

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataGroupMembers() *schema.Resource {
	return &schema.Resource{
		ReadWithoutTimeout: dataReadGroupMembers,

		DeprecationMessage: "The Bitbucket 1.0 Groups API has been permanently deprecated by Atlassian " +
			"with no replacement. This data source will be removed in a future major version of the provider. " +
			"Remove it from your configuration.",

		Schema: map[string]*schema.Schema{
			"workspace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": {
				Type:     schema.TypeString,
				Required: true,
			},
			"members": {
				Type:       schema.TypeSet,
				Elem:       &schema.Schema{Type: schema.TypeString},
				Computed:   true,
				Deprecated: "use group_members instead",
			},
			"group_members": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"username": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataReadGroupMembers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[WARN] data.bitbucket_group_members: Bitbucket 1.0 Groups API is permanently deprecated, removing from state")
	d.SetId("")
	return nil
}
