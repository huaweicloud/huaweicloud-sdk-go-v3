package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ImportApiDefinitionsV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`

	Body *ImportApiDefinitionsV2RequestBody `json:"body,omitempty" type:"multipart"`
}

func (o ImportApiDefinitionsV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ImportApiDefinitionsV2Request struct{}"
	}

	return strings.Join([]string{"ImportApiDefinitionsV2Request", string(data)}, " ")
}
