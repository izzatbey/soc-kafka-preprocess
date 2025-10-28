package preprocess

import (
	"encoding/json"
	"fmt"
	"testing"
	"preprocess-service/data"
	"github.com/tidwall/gjson"
)

// --- Helper function to unmarshal into map for readability ---
func jsonToMap(t *testing.T, raw string) map[string]interface{} {
	var m map[string]interface{}
	if err := json.Unmarshal([]byte(raw), &m); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	return m
}

// --- Test: Authentication Success Remap ---
func TestAuthSuccessRemap(t *testing.T) {
	raw := data.AuthSuccessTestData

	result := RuleAuthSuccessRemap(raw)
	if !gjson.Get(result, "destination.user").Exists() {
		t.Errorf("expected destination.user to exist, got: %s", result)
	}
	if gjson.Get(result, "data.dstuser").Exists() {
		t.Errorf("expected data.dstuser to be deleted, got: %s", result)
	}
}

// --- Test: Privilege Escalation Remap ---
func TestPrivilegeEscalationRemap(t *testing.T) {
	raw := data.PrivilegeEscalationTestData

	result := RulePrivilegeEscalationRemap(raw)
	if gjson.Get(result, "source.user").String() != "root" {
		t.Errorf("expected source.user = root, got %s", gjson.Get(result, "source.user").String())
	}
	if gjson.Get(result, "destination.user").String() != "admin" {
		t.Errorf("expected destination.user = admin, got %s", gjson.Get(result, "destination.user").String())
	}
	if gjson.Get(result, "data.srcuser").Exists() {
		t.Errorf("expected data.srcuser to be deleted")
	}
}

// --- Test: PostgreSQL Tagging ---
func TestPostgreSQLTag(t *testing.T) {
	raw := data.PostgreSQLTestData

	result := RulePostgreSQLTag(raw)
	if gjson.Get(result, "log.tag.two").String() != "postgresql" {
		t.Errorf("expected log.tag.two = postgresql, got %s", gjson.Get(result, "log.tag.two").String())
	}
}

// --- Test: Map Severity ---
func TestMapSeverity(t *testing.T) {
	cases := []struct {
		level    int
		expected string
	}{
		{1, "informational"},
		{4, "low"},
		{8, "medium"},
		{13, "high"},
		{15, "critical"},
	}

	for _, c := range cases {
		// Create test data with specific rule level
		testData := `{"rule": {"level": ` + fmt.Sprintf("%d", c.level) + `}}`
		result := RuleMapSeverity(testData)
		if gjson.Get(result, "rule.severity").String() != c.expected {
			t.Errorf("for level %d, expected severity %s, got %s", c.level, c.expected, gjson.Get(result, "rule.severity").String())
		}
	}
}
