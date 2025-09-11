package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLimitTaskRequest Request Object
type CreateLimitTaskRequest struct {

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	Body *CreateLimitTaskRequestBody `json:"body,omitempty"`
}

func (o CreateLimitTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLimitTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateLimitTaskRequest", string(data)}, " ")
}
