package preprocess

import (
	"strings"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// Rule 1: Authentication Success Remap
func RuleAuthSuccessRemap(raw string) string {
	if strings.Contains(gjson.Get(raw, "rule.groups").String(), "authentication_success") {
		fields := []string{
			"data.dstuser", 
			"data.win.eventdata.targetUserName", 
			"data.win.eventdata.remoteUserID",
		}
		for _, old := range fields {
			if val := gjson.Get(raw, old); val.Exists() {
				raw, _ = sjson.Set(raw, "destination.user", val.String())
				raw, _ = sjson.Delete(raw, old)
			}
		}
	}
	return raw
}

// Rule 2: Privilege Escalation Remap
func RulePrivilegeEscalationRemap(raw string) string {
	program := gjson.Get(raw, "predecoder.program_name").String()
	eventID := gjson.Get(raw, "data.win.system.eventID").String()

	if strings.Contains(program, "sudo") ||
		strings.Contains(eventID, "4672") ||
		strings.Contains(eventID, "4673") {

		sourceCandidates := []string{"data.srcuser", "data.win.eventdata.subjectUserName"}
		destCandidates := []string{"data.dstuser"}

		for _, old := range sourceCandidates {
			if val := gjson.Get(raw, old); val.Exists() {
				raw, _ = sjson.Set(raw, "source.user", val.String())
				raw, _ = sjson.Delete(raw, old)
			}
		}

		for _, old := range destCandidates {
			if val := gjson.Get(raw, old); val.Exists() {
				raw, _ = sjson.Set(raw, "destination.user", val.String())
				raw, _ = sjson.Delete(raw, old)
			}
		}
	}
	return raw
}

// Rule 3: PostgreSQL Tagging
func RulePostgreSQLTag(raw string) string {
	if strings.Contains(gjson.Get(raw, "agent.ip").String(), "10.80.100.17") &&
		strings.Contains(gjson.Get(raw, "location").String(), "/var/lib/pgsql/15/data/log/postgresql-") {
		raw, _ = sjson.Set(raw, "log.tag.two", "postgresql")
	}
	return raw
}

// Rule 4: Map Rule Level → Severity
func RuleMapSeverity(raw string) string {
	level := gjson.Get(raw, "rule.level")
	if !level.Exists() {
		return raw
	}

	l := level.Int()
	severity := ""

	switch {
	case l <= 2:
		severity = "informational"
	case l >= 3 && l <= 6:
		severity = "low"
	case l >= 7 && l <= 11:
		severity = "medium"
	case l >= 12 && l <= 14:
		severity = "high"
	case l >= 15:
		severity = "critical"
	default:
		return raw
	}

	raw, _ = sjson.Set(raw, "rule.severity", severity)
	return raw
}
