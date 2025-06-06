package resources

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	config "github.com/zededa/terraform-provider-zedcloud/v2/client/identity_access_management"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	testhelper "github.com/zededa/terraform-provider-zedcloud/v2/testing"
)

func TestEnterprise_Create(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping enterprise test for CI environment")
	}

	var got models.Enterprise
	var expected models.Enterprise

	// input config
	inputPath := "iam/enterprise.create.tf"
	input := testhelper.MustGetTestInput(t, inputPath)

	// expected output
	expectedPath := "iam/enterprise.create.yaml"
	testhelper.MustGetExpectedOutput(t, expectedPath, &expected)

	// terraform acceptance test case
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testhelper.CheckEnv(t) },
		CheckDestroy: testEnterpriseDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: input,
				Check: resource.ComposeTestCheckFunc(
					testEnterpriseExists("zedcloud_enterprise.test_tf_provider", &got),
					resource.TestMatchResourceAttr(
						"zedcloud_enterprise.test_tf_provider",
						"id",
						regexp.MustCompile("^[0-9A-Za-z_=-]{28}$"),
					),
					testEnterpriseAttributes(t, &got, &expected),
				),
			},
		},
	})
}

// testEnterpriseExists retrieves the Enterprise and stores it in the provided *models.DeviceConfig.
func testEnterpriseExists(resourceName string, enterpriseModel *models.Enterprise) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// retrieve the resource by name from state
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Enterprise not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Enterprise ID is not set")
		}

		// retrieve the client established in Provider configuration
		client := testProvider.Meta().(*api_client.ZedcloudAPI)

		// retrieve the ApplicationInstance by referencing its state ID for API lookup
		params := config.NewIdentityAccessManagementGetEnterpriseParams()
		params.ID = rs.Primary.ID
		response, err := client.IdentityAccessManagement.IdentityAccessManagementGetEnterprise(params, nil)
		if err != nil {
			return fmt.Errorf("could not fetch Enterprise (%s): %w", rs.Primary.ID, err)
		}

		enterprise := response.GetPayload()
		if enterprise == nil {
			return errors.New("could not get response payload in Enterprise existence test: enterprise is nil")
		}

		*enterpriseModel = *enterprise
		return nil
	}
}

// testEnterpriseAttributes verifies attributes are set correctly by Terraform
func testEnterpriseAttributes(t *testing.T, got, expected *models.Enterprise) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		ignoredFields := []string{
			"ID",
			"HubspotID",
			"SfdcID",
			"AzureSubID",
			"ParentEntpID",
			"Revision",
		}
		opts := cmpopts.IgnoreFields(models.Enterprise{}, ignoredFields...)
		if diff := cmp.Diff(*got, *expected, opts); len(diff) != 0 {
			return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), diff)
		}
		return nil
	}
}

// testEnterpriseDestroy verifies the Enterprise has been destroyed.
func testEnterpriseDestroy(s *terraform.State) error {
	// retrieve the client established in Provider configuration
	client := testProvider.Meta().(*api_client.ZedcloudAPI)

	// loop through the resources in state, verifying each Enterprise is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "zedcloud_enterprise" {
			continue
		}

		// retrieve the Enterprise by referencing its state ID for API lookup
		params := config.NewIdentityAccessManagementGetEnterpriseParams()
		params.ID = rs.Primary.ID
		response, err := client.IdentityAccessManagement.IdentityAccessManagementGetEnterprise(params, nil)
		if err == nil {
			if enterprise := response.GetPayload(); enterprise != nil && enterprise.ID == rs.Primary.ID {
				return fmt.Errorf("destroy failed, Enterprise (%s) still exists", enterprise.ID)
			}
			return nil
		}

		// if we use an http client with retries,
		// it overrrides IdentityAccessManagementGetEnterpriseNotFound error
		if strings.Contains(err.Error(), "unexpected HTTP status 404 Not Found") {
			return nil
		}

		// if the error is equivalent to 404 not found, the ApplicationInstance is destroyed
		_, ok := err.(*config.IdentityAccessManagementGetEnterpriseNotFound)
		if !ok {
			return fmt.Errorf("destroy failed, expect status code 404 for Enterprise (%s)", params.ID)
		}
	}
	return nil
}
