// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v55/common"
	oci_opsi "github.com/oracle/oci-go-sdk/v55/opsi"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	HostInsightRequiredOnlyResource = HostInsightResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_host_insight", "test_host_insight", acctest.Required, acctest.Create, hostInsightRepresentation)

	HostInsightResourceConfig = HostInsightResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_host_insight", "test_host_insight", acctest.Optional, acctest.Update, hostInsightRepresentation)

	hostInsightSingularDataSourceRepresentation = map[string]interface{}{
		"host_insight_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_opsi_host_insight.test_host_insight.id}`},
	}

	hostInsightDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"host_type":                 acctest.Representation{RepType: acctest.Optional, Create: []string{`EXTERNAL-HOST`}},
		"id":                        acctest.Representation{RepType: acctest.Optional, Create: `${oci_opsi_host_insight.test_host_insight.id}`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: []string{`ACTIVE`}},
		"status":                    acctest.Representation{RepType: acctest.Optional, Create: []string{`Enabled`}, Update: []string{`DISABLED`}},
		"filter":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: hostInsightDataSourceFilterRepresentation},
	}

	hostInsightDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_opsi_host_insight.test_host_insight.id}`}},
	}

	hostInsightRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"entity_source":       acctest.Representation{RepType: acctest.Required, Create: `MACS_MANAGED_EXTERNAL_HOST`},
		"management_agent_id": acctest.Representation{RepType: acctest.Required, Create: `${var.managed_agent_id}`},
		"status":              acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `DISABLED`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":           acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesHostInsightRepresentation},
	}

	ignoreChangesHostInsightRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	HostInsightResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: opsi/controlPlane
func TestOpsiHostInsightResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOpsiHostInsightResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	managementAgentId := utils.GetEnvSettingWithBlankDefault("managed_agent_id")
	if managementAgentId == "" {
		t.Skip("Manual install agent and set managed_agent_id to run this test")
	}
	managementAgentIdVariableStr := fmt.Sprintf("variable \"managed_agent_id\" { default = \"%s\" }\n", managementAgentId)

	resourceName := "oci_opsi_host_insight.test_host_insight"
	datasourceName := "data.oci_opsi_host_insights.test_host_insights"
	singularDatasourceName := "data.oci_opsi_host_insight.test_host_insight"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+managementAgentIdVariableStr+HostInsightResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_host_insight", "test_host_insight", acctest.Optional, acctest.Create, hostInsightRepresentation), "opsi", "hostInsight", t)

	acctest.ResourceTest(t, testAccCheckOpsiHostInsightDestroy, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + managementAgentIdVariableStr + HostInsightResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_host_insight", "test_host_insight", acctest.Optional, acctest.Create, hostInsightRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "entity_source", "MACS_MANAGED_EXTERNAL_HOST"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "host_name"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "management_agent_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "status"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + managementAgentIdVariableStr + compartmentIdUVariableStr + HostInsightResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_host_insight", "test_host_insight", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(hostInsightRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "entity_source", "MACS_MANAGED_EXTERNAL_HOST"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "host_name"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "management_agent_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "status"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + managementAgentIdVariableStr + HostInsightResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_host_insight", "test_host_insight", acctest.Optional, acctest.Update, hostInsightRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "entity_source", "MACS_MANAGED_EXTERNAL_HOST"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "host_name"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "management_agent_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "status"),
				resource.TestCheckResourceAttr(resourceName, "status", "DISABLED"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_opsi_host_insights", "test_host_insights", acctest.Optional, acctest.Update, hostInsightDataSourceRepresentation) +
				compartmentIdVariableStr + managementAgentIdVariableStr + HostInsightResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_host_insight", "test_host_insight", acctest.Optional, acctest.Update, hostInsightRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttr(datasourceName, "host_type.#", "1"),
				//resource.TestCheckResourceAttr(datasourceName, "id.#", "1"), //id is not list and it is a string So ignoring this field
				resource.TestCheckResourceAttr(datasourceName, "state.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "status.#", "1"),

				resource.TestCheckResourceAttr(datasourceName, "host_insight_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "host_insight_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_opsi_host_insight", "test_host_insight", acctest.Required, acctest.Create, hostInsightSingularDataSourceRepresentation) +
				compartmentIdVariableStr + managementAgentIdVariableStr + HostInsightResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "host_insight_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "entity_source", "MACS_MANAGED_EXTERNAL_HOST"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "host_display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "host_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "host_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "platform_name"), // Not Supported currently
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "platform_type"), // Not Supported currently
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "platform_version"), //Not Supported currently
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "processor_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttr(singularDatasourceName, "status", "DISABLED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + managementAgentIdVariableStr + HostInsightResourceConfig,
		},
		// Verify enable operation
		{
			Config: config + compartmentIdVariableStr + managementAgentIdVariableStr + HostInsightResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_host_insight", "test_host_insight", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(hostInsightRepresentation, map[string]interface{}{
						"status": acctest.Representation{RepType: acctest.Required, Update: `ENABLED`},
					})),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckOpsiHostInsightDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).OperationsInsightsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_opsi_host_insight" {
			noResourceFound = false
			request := oci_opsi.GetHostInsightRequest{}

			tmp := rs.Primary.ID
			request.HostInsightId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "opsi")

			response, err := client.GetHostInsight(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_opsi.LifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.GetLifecycleState())]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.GetLifecycleState())
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("OpsiHostInsight") {
		resource.AddTestSweepers("OpsiHostInsight", &resource.Sweeper{
			Name:         "OpsiHostInsight",
			Dependencies: acctest.DependencyGraph["hostInsight"],
			F:            sweepOpsiHostInsightResource,
		})
	}
}

func sweepOpsiHostInsightResource(compartment string) error {
	operationsInsightsClient := acctest.GetTestClients(&schema.ResourceData{}).OperationsInsightsClient()
	hostInsightIds, err := getHostInsightIds(compartment)
	if err != nil {
		return err
	}
	for _, hostInsightId := range hostInsightIds {
		if ok := acctest.SweeperDefaultResourceId[hostInsightId]; !ok {

			disableHostInsightRequest := oci_opsi.DisableHostInsightRequest{}
			disableHostInsightRequest.HostInsightId = &hostInsightId
			disableHostInsightRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "opsi")
			_, error := operationsInsightsClient.DisableHostInsight(context.Background(), disableHostInsightRequest)
			if error != nil {
				fmt.Printf("Error disabling HostInsight %s %s, It is possible that the resource is already disabled. Please verify manually \n", hostInsightId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &hostInsightId, hostInsightDisableSweepWaitCondition, time.Duration(3*time.Minute),
				hostInsightSweepResponseFetchOperation, "opsi", true)

			deleteHostInsightRequest := oci_opsi.DeleteHostInsightRequest{}

			deleteHostInsightRequest.HostInsightId = &hostInsightId

			deleteHostInsightRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "opsi")
			_, error1 := operationsInsightsClient.DeleteHostInsight(context.Background(), deleteHostInsightRequest)
			if error1 != nil {
				fmt.Printf("Error deleting HostInsight %s %s, It is possible that the resource is already deleted. Please verify manually \n", hostInsightId, error1)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &hostInsightId, hostInsightSweepWaitCondition, time.Duration(3*time.Minute),
				hostInsightSweepResponseFetchOperation, "opsi", true)
		}
	}
	return nil
}

func getHostInsightIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "HostInsightId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	operationsInsightsClient := acctest.GetTestClients(&schema.ResourceData{}).OperationsInsightsClient()

	listHostInsightsRequest := oci_opsi.ListHostInsightsRequest{}
	listHostInsightsRequest.CompartmentId = &compartmentId
	listHostInsightsRequest.LifecycleState = []oci_opsi.LifecycleStateEnum{oci_opsi.LifecycleStateActive}
	listHostInsightsResponse, err := operationsInsightsClient.ListHostInsights(context.Background(), listHostInsightsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting HostInsight list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, hostInsight := range listHostInsightsResponse.Items {
		id := *hostInsight.GetId()
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "HostInsightId", id)
	}
	return resourceIds, nil
}

func hostInsightSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if hostInsightResponse, ok := response.Response.(oci_opsi.GetHostInsightResponse); ok {
		return hostInsightResponse.GetLifecycleState() != oci_opsi.LifecycleStateDeleted
	}
	return false
}

func hostInsightSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OperationsInsightsClient().GetHostInsight(context.Background(), oci_opsi.GetHostInsightRequest{
		HostInsightId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}

func hostInsightDisableSweepWaitCondition(response common.OCIOperationResponse) bool {
	if hostInsightResponse, ok := response.Response.(oci_opsi.GetHostInsightResponse); ok {
		return hostInsightResponse.GetStatus() != oci_opsi.ResourceStatusDisabled
	}
	return false
}
