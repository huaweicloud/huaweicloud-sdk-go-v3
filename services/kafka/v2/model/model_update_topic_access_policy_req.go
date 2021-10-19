package model

import (
	"encoding/json"

	"strings"
)

type UpdateTopicAccessPolicyReq struct {
	// topic列表。

	Topics []AccessPolicyTopicEntity `json:"topics"`
}

func (o UpdateTopicAccessPolicyReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateTopicAccessPolicyReq struct{}"
	}

	return strings.Join([]string{"UpdateTopicAccessPolicyReq", string(data)}, " ")
}
