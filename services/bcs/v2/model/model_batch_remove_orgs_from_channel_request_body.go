package model

import (
	"encoding/json"

	"strings"
)

type BatchRemoveOrgsFromChannelRequestBody struct {
	// 组织名称列表

	OrgNames []string `json:"org_names"`
}

func (o BatchRemoveOrgsFromChannelRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchRemoveOrgsFromChannelRequestBody struct{}"
	}

	return strings.Join([]string{"BatchRemoveOrgsFromChannelRequestBody", string(data)}, " ")
}
