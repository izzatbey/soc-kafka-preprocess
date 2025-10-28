package preprocess

func ApplyPreprocessRules(raw string) string {
	raw = RuleAuthSuccessRemap(raw)
	raw = RulePrivilegeEscalationRemap(raw)
	raw = RulePostgreSQLTag(raw)
	raw = RuleMapSeverity(raw)
	return raw
}
