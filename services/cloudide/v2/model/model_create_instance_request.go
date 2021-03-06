package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateInstanceRequest struct {
	// 组织id（对应华为云帐号的domainId）

	OrgId string `json:"org_id"`

	Body *InstanceParam `json:"body,omitempty"`
}

func (o CreateInstanceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateInstanceRequest struct{}"
	}

	return strings.Join([]string{"CreateInstanceRequest", string(data)}, " ")
}
