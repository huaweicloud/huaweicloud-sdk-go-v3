package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProtectStatusStatistics 防护状态统计信息
type ProtectStatusStatistics struct {

	// **参数解释**: 防护类型 **取值范围**: - host：主机 - container：容器 - application：应用 - serverless：Serverless
	ProtectType *string `json:"protect_type,omitempty"`

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 未防护数 **取值范围**: 最小值0，最大值2147483647
	UnprotectedNum *int32 `json:"unprotected_num,omitempty"`

	// **参数解释**: 已防护数，只在serverless类型返回 **取值范围**: 最小值0，最大值2147483647
	ProtectedNum *int32 `json:"protected_num,omitempty"`

	// **参数解释**: 有风险数 **取值范围**: 最小值0，最大值2147483647
	RiskNum *int32 `json:"risk_num,omitempty"`

	// **参数解释**: 无风险数 **取值范围**: 最小值0，最大值2147483647
	CleanNum *int32 `json:"clean_num,omitempty"`
}

func (o ProtectStatusStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectStatusStatistics struct{}"
	}

	return strings.Join([]string{"ProtectStatusStatistics", string(data)}, " ")
}
