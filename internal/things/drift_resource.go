package things

import (
	"context"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func driftResource() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Always drift!",

		CreateContext: resourceDriftCreate,
		ReadContext:   resourceDriftRead,
		UpdateContext: resourceDriftUpdate,
		DeleteContext: resourceDriftDelete,

		Schema: map[string]*schema.Schema{
			"no_work": {
				Description: "This attribute should always drift",
				Type:        schema.TypeMap,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceDriftCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// use the meta value to retrieve your client from the things configure method
	// client := meta.(*apiClient)

	id, err := uuid.NewV7()
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(id.String())

	// write logs using the tflog package
	// see https://pkg.go.dev/github.com/hashicorp/terraform-plugin-log/tflog
	// for more information
	tflog.Trace(ctx, fmt.Sprintf("created drift with id %s", id.String()))

	return nil
}

func resourceDriftRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// use the meta value to retrieve your client from the things configure method
	// client := meta.(*apiClient)

	tflog.Debug(ctx, "setting empty drift no_work")
	forceTags := map[string]interface{}{}
	d.Set("tags", forceTags)
	return nil
}

func resourceDriftUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// use the meta value to retrieve your client from the things configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}

func resourceDriftDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// use the meta value to retrieve your client from the things configure method
	// client := meta.(*apiClient)

	return nil
}
