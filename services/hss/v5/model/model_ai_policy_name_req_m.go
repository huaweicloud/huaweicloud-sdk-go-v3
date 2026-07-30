package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AiPolicyNameReqM **参数解释**： 策略名称 **约束限制**： 必填 **取值范围**： - 0: 意图行为一致性检测 - 1: 命令执行控制 - 2: 文件访问控制 - 3: 敏感信息检测 - 4: 角色限定  **默认取值**： 不涉及
type AiPolicyNameReqM struct {
}

func (o AiPolicyNameReqM) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiPolicyNameReqM struct{}"
	}

	return strings.Join([]string{"AiPolicyNameReqM", string(data)}, " ")
}
