package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetNewBackupPolicyRequest Request Object
type SetNewBackupPolicyRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
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
