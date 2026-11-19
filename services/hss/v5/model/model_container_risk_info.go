package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ContainerRiskInfo struct {

	// **参数解释**: 基线风险数量 **取值范围**: 最小值0，最大值2147483647
	BaselineRiskNum *int32 `json:"baseline_risk_num,omitempty"`

	// **参数解释**: 漏洞总数 **取值范围**: 最小值0，最大值2147483647
	VulTotalNum *int32 `json:"vul_total_num,omitempty"`

	// **参数解释**: 紧急漏洞数量 **取值范围**: 最小值0，最大值2147483647
	CriticalVulNum *int32 `json:"critical_vul_num,omitempty"`
}

func (o ContainerRiskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContainerRiskInfo struct{}"
	}

	return strings.Join([]string{"ContainerRiskInfo", string(data)}, " ")
}
