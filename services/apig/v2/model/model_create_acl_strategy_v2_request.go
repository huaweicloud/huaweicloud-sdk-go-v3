package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateAclStrategyV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`

	Body *ApiAclCreate `json:"body,omitempty"`
}

func (o CreateAclStrategyV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateAclStrategyV2Request struct{}"
	}

	return strings.Join([]string{"CreateAclStrategyV2Request", string(data)}, " ")
}
