package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AiPolicyList struct {

	// **参数解释**: 策略ID **取值范围**: 字符长度1-20位
	PolicyId *string `json:"policy_id,omitempty"`

	// **参数解释**: 策略名称 **取值范围**： - 0: 意图行为一致性检测 - 1: 命令执行控制 - 2: 文件访问控制 - 3: 敏感信息检测 - 4: 角色限定
	PolicyName *string `json:"policy_name,omitempty"`

	// **参数解释**: 是否启用 **取值范围**: - false：否 - true：是
	Enabled *bool `json:"enabled,omitempty"`
}

func (o AiPolicyList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiPolicyList struct{}"
	}

	return strings.Join([]string{"AiPolicyList", string(data)}, " ")
}
