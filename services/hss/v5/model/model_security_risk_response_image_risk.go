package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityRiskResponseImageRisk **参数解释**： 镜像风险
type SecurityRiskResponseImageRisk struct {

	// **参数解释**： 扣分 **取值范围**: 最小值0，最大值10
	DeductScore *float32 `json:"deduct_score,omitempty"`

	// **参数解释**： 未扫描镜像数 **取值范围**: 最小值0，最大值2147483647
	UnScannedImageNum *int32 `json:"un_scanned_image_num,omitempty"`

	// **参数解释**： 风险列表 **取值范围**: 最小值0，最大值4
	RiskList *[]SecurityRiskResponseImageRiskRiskList `json:"risk_list,omitempty"`

	// **参数解释**： 风险总数 **取值范围**: 最小值0，最大值2147483647
	TotalRiskNum *int32 `json:"total_risk_num,omitempty"`
}

func (o SecurityRiskResponseImageRisk) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityRiskResponseImageRisk struct{}"
	}

	return strings.Join([]string{"SecurityRiskResponseImageRisk", string(data)}, " ")
}
