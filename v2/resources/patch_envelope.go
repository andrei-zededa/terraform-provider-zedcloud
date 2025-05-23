package resources

import (
	"context"
	"errors"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	config "github.com/zededa/terraform-provider-zedcloud/v2/client/patch_envelope"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
)

func PatchEnvelopeResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: CreatePatchEnvelope,
		ReadContext:   ReadPatchEnvelope,
		UpdateContext: UpdatePatchEnvelope,
		DeleteContext: DeletePatchEnvelope,
		Schema:        zschema.PatchEnvelopeSchema(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func PatchEnvelopeDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: ReadPatchEnvelope,
		Schema:      zschema.PatchEnvelopeSchema(),
	}
}

func CreatePatchEnvelope(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.PatchEnvelopeModel(d)
	params := config.NewPatchEnvelopeConfigurationCreatePatchEnvelopeParams()
	params.SetBody(model)

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.PatchEnvelope.PatchEnvelopeConfigurationCreatePatchEnvelope(params, nil)
	if err != nil {
		log.Printf("[TRACE] patch envelope create error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("patch envelope create error: %s", err)...)
		return diags
	}

	responseData := resp.GetPayload()
	if responseData != nil && len(responseData.Error) > 0 {
		for _, err := range responseData.Error {
			// FIXME: zedcloud api returns a response that contains and error even in case of success.
			// remove this code once it is fixed on API side.
			if err.ErrorCode != nil && *err.ErrorCode == models.ErrorCodeSuccess {
				continue
			}
			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
		}
		if diags.HasError() {
			return diags
		}
	}

	// note, we need to set the ID in any case, even GetByName endpoint seems to require the ID
	// but doesn't return any error if it's not set.
	d.SetId(responseData.ObjectID)

	// the zedcloud API does not return the partially updated object but a custom response.
	// thus, we need to fetch the object and populate the state.
	if diags := readPatchEnvelopeByName(ctx, d, m); diags.HasError() {
		return diags
	}

	return diags
}

func ReadPatchEnvelope(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if _, isSet := d.GetOk("name"); isSet {
		return readPatchEnvelopeByName(ctx, d, m)
	}
	return readPatchEnvelopeByID(ctx, d, m)
}

func readPatchEnvelopeByID(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := config.NewPatchEnvelopeConfigurationGetPatchEnvelopeByIDParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	id, isSet := d.GetOk("id")
	if isSet {
		params.ID = id.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.PatchEnvelope.PatchEnvelopeConfigurationGetPatchEnvelopeByID(params, nil)
	if err != nil {
		log.Printf("[TRACE] patch envelope read error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("patch envelope read error: %s", err)...)
		return diags
	}

	patchEnv := resp.GetPayload()
	zschema.SetPatchEnvelopeResourceData(d, patchEnv)
	d.SetId(patchEnv.ID)

	return diags

}

func readPatchEnvelopeByName(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := config.NewPatchEnvelopeConfigurationGetPatchEnvelopeByNameParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	nameVal, nameIsSet := d.GetOk("name")
	if nameIsSet {
		params.Name = nameVal.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: name")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.PatchEnvelope.PatchEnvelopeConfigurationGetPatchEnvelopeByName(params, nil)
	if err != nil {
		log.Printf("[TRACE] patch envelope read error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("patch envelope read error: %s", err)...)
		return diags
	}

	patchEnv := resp.GetPayload()
	zschema.SetPatchEnvelopeResourceData(d, patchEnv)
	d.SetId(patchEnv.ID)

	return diags

}

func UpdatePatchEnvelope(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := config.NewPatchEnvelopeConfigurationUpdatePatchEnvelopeParams()
	model := zschema.PatchEnvelopeModel(d)

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	params.SetBody(model)

	idVal, idIsSet := d.GetOk("id")
	if idIsSet {
		id, _ := idVal.(string)
		params.ID = id
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.PatchEnvelope.PatchEnvelopeConfigurationUpdatePatchEnvelope(params, nil)
	if err != nil {
		log.Printf("[TRACE] patch envelope update error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("patch envelope update error: %s", err)...)
		return diags
	}

	responseData := resp.GetPayload()
	if responseData != nil && len(responseData.Error) > 0 {
		for _, err := range responseData.Error {
			// FIXME: zedcloud api returns a response that contains and error even in case of success.
			// remove this code once it is fixed on API side.
			if err.ErrorCode != nil && *err.ErrorCode == models.ErrorCodeSuccess {
				continue
			}
			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
		}
		if diags.HasError() {
			return diags
		}
	}

	// note, we need to set the ID in any case, even GetByName endpoint seems to require items
	// but doesn't return any error if it's not set.
	d.SetId(responseData.ObjectID)

	// the zedcloud API does not return the partially updated object but a custom response.
	// thus, we need to fetch the object and populate the state.
	if diags := readPatchEnvelopeByName(ctx, d, m); diags.HasError() {
		return diags
	}

	return diags
}

func DeletePatchEnvelope(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := config.NewPatchEnvelopeConfigurationDeletePatchEnvelopeParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	idVal, idIsSet := d.GetOk("id")
	if idIsSet {
		id, _ := idVal.(string)
		params.ID = id
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)

	_, err := client.PatchEnvelope.PatchEnvelopeConfigurationDeletePatchEnvelope(params, nil)
	if err != nil {
		log.Printf("[TRACE] patch envelope delete error: %s", spew.Sdump(err))
		if ds, ok := ZsrvResponderToDiags(err); ok {
			diags = append(diags, ds...)
			return diags
		}

		diags = append(diags, diag.Errorf("patch envelope delete error: %s", err)...)
		return diags
	}

	d.SetId("")
	return diags
}
