package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func TagLevelSettingsModel(d *schema.ResourceData) *models.TagLevelSettings {
	var flowLogTransmission *models.NetworkInstanceFlowLogTransmission // NetworkInstanceFlowLogTransmission
	flowLogTransmissionInterface, flowLogTransmissionIsSet := d.GetOk("flow_log_transmission")
	if flowLogTransmissionIsSet {
		flowLogTransmissionModel := flowLogTransmissionInterface.(string)
		flowLogTransmission = models.NewNetworkInstanceFlowLogTransmission(models.NetworkInstanceFlowLogTransmission(flowLogTransmissionModel))
	}
	var interfaceOrdering *models.InterfaceOrdering // InterfaceOrdering
	interfaceOrderingInterface, interfaceOrderingIsSet := d.GetOk("interface_ordering")
	// TODO: It seems that GetOk returns `interfaceOrderingIsSet == true`
	// even if `interface_ordering` is not actually present in the config,
	// check with something like:
	//	log.Printf("[TRACE] project interface ordering debug: %s", spew.Sdump(interfaceOrdering))
	//	log.Printf("[TRACE] project interface ordering debug: %s", spew.Sdump(interfaceOrderingIsSet))
	//	log.Printf("[TRACE] project interface ordering debug: %s", spew.Sdump(interfaceOrderingInterface))
	if interfaceOrderingIsSet {
		interfaceOrderingModel := interfaceOrderingInterface.(string)
		// TODO: If we set `interfaceOrdering` here, even if it's an
		// empty string because it's not actually part of the config it
		// will cause errors with older versions of Zedcloud which do
		// not support it. If we don't set it then it will remain `nil`
		// and won't be sent due to `omitempty`.
		if len(interfaceOrderingModel) > 0 {
			interfaceOrdering = models.NewInterfaceOrdering(models.InterfaceOrdering(interfaceOrderingModel))
		}
	}
	return &models.TagLevelSettings{
		FlowLogTransmission: flowLogTransmission,
		InterfaceOrdering:   interfaceOrdering,
	}
}

func TagLevelSettingsModelFromMap(m map[string]interface{}) *models.TagLevelSettings {
	var flowLogTransmission *models.NetworkInstanceFlowLogTransmission // NetworkInstanceFlowLogTransmission
	flowLogTransmissionInterface, flowLogTransmissionIsSet := m["flow_log_transmission"]
	if flowLogTransmissionIsSet {
		flowLogTransmissionModel := flowLogTransmissionInterface.(string)
		flowLogTransmission = models.NewNetworkInstanceFlowLogTransmission(models.NetworkInstanceFlowLogTransmission(flowLogTransmissionModel))
	}
	var interfaceOrdering *models.InterfaceOrdering // InterfaceOrdering
	interfaceOrderingInterface, interfaceOrderingIsSet := m["interface_ordering"]
	// TODO: Not tested if here we have a similar problem with `interfaceOrderingIsSet == true`
	// when it's not actually part of the config like for `GetOk` above.
	if interfaceOrderingIsSet {
		interfaceOrderingModel := interfaceOrderingInterface.(string)
		// TODO: If we set `interfaceOrdering` here, even if it's an
		// empty string it will cause errors with older versions of
		// Zedcloud which do not support it. If we don't set it then it
		// will remain `nil` and won't be sent due to `omitempty`.
		if len(interfaceOrderingModel) > 0 {
			interfaceOrdering = models.NewInterfaceOrdering(models.InterfaceOrdering(interfaceOrderingModel))
		}
	}
	return &models.TagLevelSettings{
		FlowLogTransmission: flowLogTransmission,
		InterfaceOrdering:   interfaceOrdering,
	}
}

func SetTagLevelSettingsResourceData(d *schema.ResourceData, m *models.TagLevelSettings) {
	d.Set("flow_log_transmission", m.FlowLogTransmission)
	d.Set("interface_ordering", m.InterfaceOrdering)
}

func SetTagLevelSettingsSubResourceData(m []*models.TagLevelSettings) (d []*map[string]interface{}) {
	for _, TagLevelSettingsModel := range m {
		if TagLevelSettingsModel != nil {
			properties := make(map[string]interface{})
			properties["flow_log_transmission"] = TagLevelSettingsModel.FlowLogTransmission
			properties["interface_ordering"] = TagLevelSettingsModel.InterfaceOrdering
			d = append(d, &properties)
		}
	}
	return
}

func TagLevelSettingsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"flow_log_transmission": {
			Description: `Flow log transmission setting for the network instances`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"interface_ordering": {
			Description: `interface ordering for app instances`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetTagLevelSettingsPropertyFields() (t []string) {
	return []string{
		"flow_log_transmission",
		"interface_ordering",
	}
}
