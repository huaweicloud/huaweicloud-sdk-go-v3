package model

import (
	"encoding/json"

	"strings"
)

type ApiVersionInfo struct {
	// API版本的编号

	VersionId *string `json:"version_id,omitempty"`
}

func (o ApiVersionInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ApiVersionInfo struct{}"
	}

	return strings.Join([]string{"ApiVersionInfo", string(data)}, " ")
}
