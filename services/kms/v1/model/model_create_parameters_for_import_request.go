package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateParametersForImportRequest struct {
	// API版本号

	VersionId string `json:"version_id"`

	Body *GetParametersForImportRequestBody `json:"body,omitempty"`
}

func (o CreateParametersForImportRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateParametersForImportRequest struct{}"
	}

	return strings.Join([]string{"CreateParametersForImportRequest", string(data)}, " ")
}
