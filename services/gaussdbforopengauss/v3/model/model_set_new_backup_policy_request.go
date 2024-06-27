package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetNewBackupPolicyRequest Request Object
type SetNewBackupPolicyRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	Body *SetNewBackupPolicyRequestBody `json:"body,omitempty"`
}

func (o SetNewBackupPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetNewBackupPolicyRequest struct{}"
	}

	return strings.Join([]string{"SetNewBackupPolicyRequest", string(data)}, " ")
}
