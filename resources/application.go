// // Code generated by go-swagger; DO NOT EDIT.
package resources

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api_client "github.com/zededa/terraform-provider/client"
	config "github.com/zededa/terraform-provider/client/edge_application_configuration"
	"github.com/zededa/terraform-provider/models"
	zschema "github.com/zededa/terraform-provider/schemas"
)

func ApplicationResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: CreateApplication,
		ReadContext:   ReadApplicationByName,
		UpdateContext: UpdateApplication,
		DeleteContext: DeleteApplication,
		Schema:        zschema.Application(),
	}
}

func ApplicationDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: ReadApplicationByName,
		Schema:      zschema.Application(),
	}
}

func CreateApplication(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.ApplicationModel(d)

	// we support reading the manifest from a JSON file for legacy reasons
	if path, ok := d.GetOk("manifest_file"); ok {
		manifest, err := readManifestFromFile(path.(string))
		if err != nil {
			diags = append(diags, diag.FromErr(err)...)
			return diags
		}
		model.ManifestJSON = manifest
		log.Println("[WARN] note, diffs for manifest files are suppressed and cannot be considered by the terraform plan command!")
	}

	// the ac_kind depends on the app_type set in the manifest.
	// this should be read-only and be computed in the backend but is not, so we have to compute it in the client.
	acKind, err := getAcKind(model.ManifestJSON.AppType)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}
	model.ManifestJSON.AcKind = acKind

	params := config.CreateParams()
	params.SetBody(model)

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.Application.Create(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
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
	if errs := ReadApplicationByName(ctx, d, m); err != nil {
		return append(diags, errs...)
	}

	return diags
}

func ReadApplicationByName(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := config.GetByNameParams()

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

	resp, err := client.Application.GetByName(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		return append(diags, diag.Errorf("unexpected: %s", err)...)
	}

	app := resp.GetPayload()
	zschema.SetApplicationResourceData(d, app)
	d.SetId(app.ID)

	return diags
}

func UpdateApplication(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.ApplicationModel(d)

	// we support reading the manifest from a JSON file for legacy reasons
	if path, ok := d.GetOk("manifest_file"); ok {
		manifest, err := readManifestFromFile(path.(string))
		if err != nil {
			diags = append(diags, diag.FromErr(err)...)
			return diags
		}
		model.ManifestJSON = manifest
		log.Println("[WARN] note, diffs for manifest files are suppressed and cannot be considered by the terraform plan command!")
	}

	// the ac_kind depends on the app_type set in the manifest.
	// this should be read-only and be computed in the backend but is not, so we have to compute it in the client.
	acKind, err := getAcKind(model.ManifestJSON.AppType)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}
	model.ManifestJSON.AcKind = acKind

	params := config.UpdateParams()
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

	// makes a bulk update for all properties that were changed
	client := m.(*api_client.ZedcloudAPI)
	resp, err := client.Application.Update(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		return append(diags, diag.Errorf("unexpected: %s", err)...)
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
	if errs := ReadApplicationByName(ctx, d, m); err != nil {
		return append(diags, errs...)
	}

	return diags
}

func DeleteApplication(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := config.DeleteParams()

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

	resp, err := client.Application.Delete(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return diags
	}

	d.SetId("")
	return diags
}

func readManifestFromFile(path string) (*models.VMManifest, error) {
	jsonData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read manifest file (path=%s): %w", path, err)
	}
	var manifest models.VMManifest
	err = json.Unmarshal([]byte(jsonData), &manifest)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal manifest file (path=%s): %w", path, err)
	}
	return &manifest, nil
}

func getAcKind(appType *models.AppType) (*string, error) {
	if appType == nil {
		return nil, nil
	}
	switch *appType {
	case models.AppTypeAPPTYPEVM:
		acKind := "VMManifest"
		return &acKind, nil
	case models.AppTypeAPPTYPECONTAINER:
		acKind := "PodManifest"
		return &acKind, nil
	case models.AppTypeAPPTYPEMODULE:
		acKind := "ModuleManifest"
		return &acKind, nil
	default:
		return nil, fmt.Errorf("could not determine ac_kind, unsupported app_type (%s)", *appType)
	}
}
