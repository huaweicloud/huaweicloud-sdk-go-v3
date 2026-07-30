package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAiPolicyDetailRequest Request Object
type ListAiPolicyDetailRequest struct {

	// **参数解释**： 策略组ID **约束限制**： 必填 **取值范围**： 最小值0，最大值9223372036854775807 **默认取值**： 不涉及
	PolicyId string `json:"policy_id"`

	// **参数解释**： 策略名称 **约束限制**： 必填 **取值范围**： - 0: 意图行为一致性检测 - 1: 命令执行控制 - 2: 文件访问控制 - 3: 敏感信息检测 - 4: 角色限定  **默认取值**： 不涉及
	PolicyName string `json:"policy_name"`
}

func (o ListAiPolicyDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAiPolicyDetailRequest struct{}"
	}

	return strings.Join([]string{"ListAiPolicyDetailRequest", string(data)}, " ")
}
