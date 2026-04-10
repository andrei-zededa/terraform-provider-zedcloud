package schemas

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestEdgeNodeClusterConfigModelFromMap_ExampleJSON(t *testing.T) {
	raw, err := os.ReadFile("edge_node_example_for_edge_node_cluster_config_test.json")
	if err != nil {
		t.Fatalf("failed to read test JSON: %v", err)
	}

	var edgeNode map[string]interface{}
	if err := json.Unmarshal(raw, &edgeNode); err != nil {
		t.Fatalf("failed to unmarshal JSON: %v", err)
	}

	cluster, ok := edgeNode["edgeNodeCluster"].(map[string]interface{})
	if !ok {
		t.Fatal("edgeNodeCluster field missing or not an object")
	}

	// Remap camelCase JSON keys to the snake_case keys expected by EdgeNodeClusterConfigModelFromMap.
	m := map[string]interface{}{
		"cluster_prefix":     cluster["clusterPrefix"],
		"id":                 cluster["id"],
		"is_master":          cluster["isMaster"],
		"manifest":           cluster["manifest"],
		"name":               cluster["name"],
		"project_id":         cluster["projectId"],
		"seed_node_id":       cluster["seedNodeId"],
		"seed_node_ip":       cluster["seedNodeIp"],
		"tags":               cluster["tags"],
		"tie_breaker_node_id": cluster["tieBreakerNodeId"],
		"token":              cluster["token"],
	}

	// The function does bare type assertions and will panic on unexpected types.
	result := EdgeNodeClusterConfigModelFromMap(m)

	if result.ID != "35ca7d70-a4cd-4717-b141-de2f49941a50" {
		t.Errorf("unexpected ID: %s", result.ID)
	}
	if result.Name != "010-Cluster" {
		t.Errorf("unexpected Name: %s", result.Name)
	}
	if !result.IsMaster {
		t.Error("expected IsMaster to be true")
	}
	if result.ClusterPrefix != "10.244.244.3/28" {
		t.Errorf("unexpected ClusterPrefix: %s", result.ClusterPrefix)
	}
	if result.Token != "TZKvn24DZgAGOqMlzM+bARPglcvoaNtDG7G4Zt3bjhM=" {
		t.Errorf("unexpected Token: %s", result.Token)
	}
}

func TestEdgeNodeClusterConfigModel_ExampleJSON(t *testing.T) {
	raw, err := os.ReadFile("edge_node_example_for_edge_node_cluster_config_test.json")
	if err != nil {
		t.Fatalf("failed to read test JSON: %v", err)
	}

	var edgeNode map[string]interface{}
	if err := json.Unmarshal(raw, &edgeNode); err != nil {
		t.Fatalf("failed to unmarshal JSON: %v", err)
	}

	cluster, ok := edgeNode["edgeNodeCluster"].(map[string]interface{})
	if !ok {
		t.Fatal("edgeNodeCluster field missing or not an object")
	}

	values := map[string]interface{}{
		"cluster_prefix":      cluster["clusterPrefix"],
		"id":                  cluster["id"],
		"is_master":           cluster["isMaster"],
		"manifest":            cluster["manifest"],
		"name":                cluster["name"],
		"project_id":          cluster["projectId"],
		"seed_node_id":        cluster["seedNodeId"],
		"seed_node_ip":        cluster["seedNodeIp"],
		"tags":                map[string]interface{}{},
		"tie_breaker_node_id": cluster["tieBreakerNodeId"],
		"token":               cluster["token"],
	}

	rd := schema.TestResourceDataRaw(t, EdgeNodeClusterConfigSchema(), values)

	result := EdgeNodeClusterConfigModel(rd)

	if result.ID != "35ca7d70-a4cd-4717-b141-de2f49941a50" {
		t.Errorf("unexpected ID: %s", result.ID)
	}
	if result.Name != "010-Cluster" {
		t.Errorf("unexpected Name: %s", result.Name)
	}
	if !result.IsMaster {
		t.Error("expected IsMaster to be true")
	}
}
