package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CopyAiPolicyGroupRequestInfo struct {

	// **参数解释**： 策略组ID **约束限制**： 必填 **取值范围**： 字符长度1-20位 **默认取值**： 不涉及
	GroupId string `json:"group_id"`

	// **参数解释**： 策略组名称 **约束限制**： 必填 **取值范围**： 字符长度1-128位 **默认取值**： 不涉及
	GroupName string `json:"group_name"`

	// **参数解释**： 策略描述 **约束限制**： 不涉及 **取值范围**： 字符长度1-256位 **默认取值**： 不涉及
	Description *string `json:"description,omitempty"`
}

func (o CopyAiPolicyGroupRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyAiPolicyGroupRequestInfo struct{}"
	}

	return strings.Join([]string{"CopyAiPolicyGroupRequestInfo", string(data)}, " ")
}
