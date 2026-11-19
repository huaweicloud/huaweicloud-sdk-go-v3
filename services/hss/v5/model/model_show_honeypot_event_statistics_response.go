package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHoneypotEventStatisticsResponse Response Object
type ShowHoneypotEventStatisticsResponse struct {

	// **参数解释**： 存在风险的服务器数量 **取值范围**： 最小值0，最大值2147483647
	RiskyHostNum *int32 `json:"risky_host_num,omitempty"`

	// **参数解释**： 防护服务器数量 **取值范围**： 最小值0，最大值2147483647
	ProtectionHostNum *int32 `json:"protection_host_num,omitempty"`

	// **参数解释**： 致命风险数量 **取值范围**： 最小值0，最大值2147483647
	CriticalRiskyNum *int32 `json:"critical_risky_num,omitempty"`

	// **参数解释**： 高危风险数量 **取值范围**： 最小值0，最大值2147483647
	HighRiskyNum *int32 `json:"high_risky_num,omitempty"`

	// **参数解释**： 中危风险数量 **取值范围**： 最小值0，最大值2147483647
	MediumRiskyNum *int32 `json:"medium_risky_num,omitempty"`

	// **参数解释**： 低危风险数量 **取值范围**： 最小值0，最大值2147483647
	LowRiskyNum    *int32 `json:"low_risky_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowHoneypotEventStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHoneypotEventStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowHoneypotEventStatisticsResponse", string(data)}, " ")
}
