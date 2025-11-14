package preprocess

import (
	"log"
	"strings"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func ApplyPreprocessRules(raw string, logTag string) string {
	raw, _ = sjson.Set(raw, "log.tag", logTag)
	raw = DropMessage(raw)
	raw = RuleAuthSuccessRemap(raw)
	raw = RulePrivilegeEscalationRemap(raw)
	raw = RulePostgreSQLTag(raw)
	raw = RuleMapSeverity(raw)
	return raw
}

func DropMessage(raw string) string {
	paths := []string{
		"predecoder.program_name",
	}
	ignoreList := []string{
		"fluent-bit",
		"opensearch-dashboards",
		"backup",
		"localcli",
		"CROND",
	}

	for _, path := range paths {
		prog := strings.ToLower(gjson.Get(raw, path).String())
		if prog == "" {
			continue
		}

		for _, ignore := range ignoreList {
			if strings.Contains(prog, ignore) {
				log.Printf("Removing Message")
				return ""
			}
		}
	}
	return raw
}
