package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SetBackupPolicyRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
	// 实例ID，严格匹配UUID规则。

	InstanceId string `json:"instance_id"`

	Body *SetBackupPolicyRequestBody `json:"body,omitempty"`
}

func (o SetBackupPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetBackupPolicyRequest struct{}"
	}

	return strings.Join([]string{"SetBackupPolicyRequest", string(data)}, " ")
}
