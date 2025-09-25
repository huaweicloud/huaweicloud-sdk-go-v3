package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityRiskResponse Response Object
type ListSecurityRiskResponse struct {
	AlarmRisk *AlarmRiskInfo `json:"alarm_risk,omitempty"`

	BaselineRisk *SecurityRiskResponseBaselineRisk `json:"baseline_risk,omitempty"`

	AssetRisk *SecurityRiskResponseAssetRisk `json:"asset_risk,omitempty"`

	SecurityProtectRisk *SecurityRiskResponseSecurityProtectRisk `json:"security_protect_risk,omitempty"`

	VulRisk *SecurityRiskResponseVulRisk `json:"vul_risk,omitempty"`

	ImageRisk *SecurityRiskResponseImageRisk `json:"image_risk,omitempty"`

	// **参数解释**： 总风险数量 **取值范围**: 最小值0，最大值2147483647
	TotalRiskNum   *int32 `json:"total_risk_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSecurityRiskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityRiskResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityRiskResponse", string(data)}, " ")
}
