package bitbucket

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type UserGroupMembership struct {
	UUID string `json:"uuid,omitempty"`
}

func resourceGroupMembership() *schema.Resource {
	return &schema.Resource{
		CreateWithoutTimeout: resourceGroupMembershipsPut,
		ReadWithoutTimeout:   resourceGroupMembershipsRead,
		DeleteWithoutTimeout: resourceGroupMembershipsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		DeprecationMessage: "The Bitbucket 1.0 Groups API has been permanently deprecated by Atlassian " +
			"with no replacement. This resource will be removed in a future major version of the provider. " +
			"Remove it from your configuration and run `terraform state rm <address>` to drop it from state.",

		Schema: map[string]*schema.Schema{
			"workspace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"group_slug": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"slug": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGroupMembershipsPut(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return diag.Errorf("bitbucket_group_membership: the Bitbucket 1.0 Groups API has been permanently " +
		"deprecated. New group memberships cannot be created. Remove this resource from your configuration.")
}

func resourceGroupMembershipsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[WARN] Group Membership (%s): Bitbucket 1.0 Groups API is permanently deprecated, removing from state", d.Id())
	d.SetId("")
	return nil
}

func resourceGroupMembershipsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func groupMemberId(id string) (string, string, string, error) {
	parts := strings.Split(id, "/")

	if len(parts) != 3 {
		return "", "", "", fmt.Errorf("unexpected format of ID (%q), expected WORKSPACE-ID/GROUP-SLUG-ID/MEMBER-UUID", id)
	}

	return parts[0], parts[1], parts[2], nil
}
