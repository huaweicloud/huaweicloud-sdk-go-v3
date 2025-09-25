package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityRiskResponseSecurityProtectRisk **参数解释**： 安全防护风险
type SecurityRiskResponseSecurityProtectRisk struct {

	// **参数解释**： 未开主机安全防护主机数 **取值范围**: 最小值0，最大值2147483647
	UnOpenProtectionHostNum *int32 `json:"un_open_protection_host_num,omitempty"`

	// **参数解释**： 扣分 **取值范围**: 最小值0，最大值10
	DeductScore *float32 `json:"deduct_score,omitempty"`

	// **参数解释**： 风险总数 **取值范围**: 最小值0，最大值2147483647
	TotalRiskNum *int32 `json:"total_risk_num,omitempty"`
}

func (o SecurityRiskResponseSecurityProtectRisk) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityRiskResponseSecurityProtectRisk struct{}"
	}

	return strings.Join([]string{"SecurityRiskResponseSecurityProtectRisk", string(data)}, " ")
}
