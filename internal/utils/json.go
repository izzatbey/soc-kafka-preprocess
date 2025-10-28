package utils

import "encoding/json"

func PrettyPrintJSON(v interface{}) string {
	out, _ := json.MarshalIndent(v, "", "  ")
	return string(out)
}
