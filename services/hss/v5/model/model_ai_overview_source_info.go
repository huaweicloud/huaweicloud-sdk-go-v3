package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AiOverviewSourceInfo AI Defender智能体防护信息（type=source时返回）
type AiOverviewSourceInfo struct {

	// **参数解释**： 智能体总数 **取值范围**： 最小值0，最大值2000000
	TotalAgents *int32 `json:"total_agents,omitempty"`

	// **参数解释**： 已防护智能体数量 **取值范围**： 最小值0，最大值2000000
	ProtectedAgents *int32 `json:"protected_agents,omitempty"`

	// **参数解释**： 未防护智能体数量 **取值范围**： 最小值0，最大值2000000
	UnprotectedAgents *int32 `json:"unprotected_agents,omitempty"`
}

func (o AiOverviewSourceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiOverviewSourceInfo struct{}"
	}

	return strings.Join([]string{"AiOverviewSourceInfo", string(data)}, " ")
}
