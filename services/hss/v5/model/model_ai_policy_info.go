package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AiPolicyInfo struct {

	// **参数解释**: 策略ID **取值范围**: 字符长度1-20位
	PolicyId *string `json:"policy_id,omitempty"`

	// **参数解释**: 策略名称 **取值范围**： - 0: 意图行为一致性检测 - 1: 命令执行控制 - 2: 文件访问控制 - 3: 敏感信息检测 - 4: 角色限定
	PolicyName *string `json:"policy_name,omitempty"`

	// **参数解释**: 是否启用 **取值范围**: - false：否 - true：是
	Enabled *bool `json:"enabled,omitempty"`

	// **参数解释**： 策略组ID **取值范围**： 字符长度1-20位
	PolicyGroupId *string `json:"policy_group_id,omitempty"`

	// **参数解释**: 策略详情 **取值范围**: 字符长度0-65535位
	Content *string `json:"content,omitempty"`

	// **参数解释**: 策略描述 **取值范围**: 字符长度1-256位
	Description *string `json:"description,omitempty"`

	// **参数解释**： 创建时间 **取值范围**： 最小值0，最大值9223372036854775807
	CreateTime *int64 `json:"create_time,omitempty"`

	// **参数解释**： 创建时间 **取值范围**： 最小值0，最大值9223372036854775807
	UpdateTime *int64 `json:"update_time,omitempty"`
}

func (o AiPolicyInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiPolicyInfo struct{}"
	}

	return strings.Join([]string{"AiPolicyInfo", string(data)}, " ")
}
