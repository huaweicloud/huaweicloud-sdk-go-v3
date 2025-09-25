package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RiskListInfo struct {

	// **参数解释**： 风险级别 **取值范围**: 字符长度0-32位
	Severity *string `json:"severity,omitempty"`

	// **参数解释**： 风险数量 **取值范围**: 最小值0，最大值2147483647
	RiskNum *int32 `json:"risk_num,omitempty"`

	// **参数解释**： 受影响资产数量 **取值范围**: 最小值0，最大值2147483647
	EffectedHostNum *int32 `json:"effected_host_num,omitempty"`
}

func (o RiskListInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RiskListInfo struct{}"
	}

	return strings.Join([]string{"RiskListInfo", string(data)}, " ")
}
