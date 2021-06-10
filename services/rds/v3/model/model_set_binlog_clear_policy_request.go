package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type SetBinlogClearPolicyRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
	// 实例ID。

	InstanceId string `json:"instance_id"`

	Body *BinlogClearPolicyRequestBody `json:"body,omitempty"`
}

func (o SetBinlogClearPolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SetBinlogClearPolicyRequest struct{}"
	}

	return strings.Join([]string{"SetBinlogClearPolicyRequest", string(data)}, " ")
}
