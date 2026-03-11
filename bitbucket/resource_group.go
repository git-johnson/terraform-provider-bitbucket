package bitbucket

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

type UserGroup struct {
	Name                    string `json:"name,omitempty"`
	Slug                    string `json:"slug,omitempty"`
	AutoAdd                 bool   `json:"auto_add,omitempty"`
	Permission              string `json:"permission,omitempty"`
	EmailForwardingDisabled bool   `json:"email_forwarding_disabled,omitempty"`
}

func resourceGroup() *schema.Resource {
	return &schema.Resource{
		CreateWithoutTimeout: resourceGroupsCreate,
		ReadWithoutTimeout:   resourceGroupsRead,
		UpdateWithoutTimeout: resourceGroupsUpdate,
		DeleteWithoutTimeout: resourceGroupsDelete,
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
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"auto_add": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"permission": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"read", "write", "admin"}, false),
			},
			"email_forwarding_disabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func resourceGroupsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return diag.Errorf("bitbucket_group: the Bitbucket 1.0 Groups API has been permanently " +
		"deprecated. New groups cannot be created. Remove this resource from your configuration.")
}

func resourceGroupsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[WARN] Group (%s): Bitbucket 1.0 Groups API is permanently deprecated, removing from state", d.Id())
	d.SetId("")
	return nil
}

func resourceGroupsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(Clients).httpClient

	group := expandGroup(d)
	log.Printf("[DEBUG] Group Request: %#v", group)
	bytedata, err := json.Marshal(group)

	if err != nil {
		return diag.FromErr(err)
	}

	_, err = client.Put(fmt.Sprintf("1.0/groups/%s/%s/",
		d.Get("workspace").(string), d.Get("slug").(string)), bytes.NewBuffer(bytedata))

	if err != nil {
		return diag.FromErr(err)
	}

	return resourceGroupsRead(ctx, d, m)
}

func resourceGroupsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func expandGroup(d *schema.ResourceData) *UserGroup {
	group := &UserGroup{
		Name: d.Get("name").(string),
	}

	if v, ok := d.GetOk("auto_add"); ok {
		group.AutoAdd = v.(bool)
	}

	if v, ok := d.GetOk("permission"); ok && v.(string) != "" {
		group.Permission = v.(string)
	}

	if v, ok := d.GetOk("email_forwarding_disabled"); ok {
		group.EmailForwardingDisabled = v.(bool)
	}

	return group
}
