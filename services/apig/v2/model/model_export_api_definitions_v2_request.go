package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ExportApiDefinitionsV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`

	Body *SwaggerReq `json:"body,omitempty"`
}

func (o ExportApiDefinitionsV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ExportApiDefinitionsV2Request struct{}"
	}

	return strings.Join([]string{"ExportApiDefinitionsV2Request", string(data)}, " ")
}
