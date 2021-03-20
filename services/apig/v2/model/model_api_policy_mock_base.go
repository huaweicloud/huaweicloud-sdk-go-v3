package model

import (
	"encoding/json"

	"strings"
)

type ApiPolicyMockBase struct {
	// 返回结果

	ResultContent *string `json:"result_content,omitempty"`
}

func (o ApiPolicyMockBase) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ApiPolicyMockBase struct{}"
	}

	return strings.Join([]string{"ApiPolicyMockBase", string(data)}, " ")
}
