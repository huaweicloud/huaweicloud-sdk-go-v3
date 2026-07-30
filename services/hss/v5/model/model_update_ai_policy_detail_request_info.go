package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateAiPolicyDetailRequestInfo struct {

	// **参数解释**: 策略ID **约束限制**: 必填 **取值范围**: 字符长度1-20位 **默认取值**: 不涉及
	PolicyId string `json:"policy_id"`

	// **参数解释**： 策略名称 **约束限制**： 必填 **取值范围**： - 0: 意图行为一致性检测 - 1: 命令执行控制 - 2: 文件访问控制 - 3: 敏感信息检测 - 4: 角色限定  **默认取值**： 不涉及
	PolicyName string `json:"policy_name"`

	// **参数解释**： 策略详情 **约束限制**： 必填 **取值范围**： 字符长度1-65535位 **默认取值**： 不涉及
	Content string `json:"content"`
}

func (o UpdateAiPolicyDetailRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAiPolicyDetailRequestInfo struct{}"
	}

	return strings.Join([]string{"UpdateAiPolicyDetailRequestInfo", string(data)}, " ")
}
