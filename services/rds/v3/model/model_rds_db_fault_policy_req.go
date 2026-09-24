package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RdsDbFaultPolicyReq 设置内核故障的处理策略请求体。
type RdsDbFaultPolicyReq struct {

	// **参数解释**：  内核故障的处理策略。  **约束限制**：  不涉及。  **取值范围**：  - repairFirst：优先修复 - failoverFirst：优先切换  **默认取值**：  不涉及。
	DbPolicy string `json:"db_policy"`
}

func (o RdsDbFaultPolicyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdsDbFaultPolicyReq struct{}"
	}

	return strings.Join([]string{"RdsDbFaultPolicyReq", string(data)}, " ")
}
