package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTrialsStatisticsResponse Response Object
type ShowTrialsStatisticsResponse struct {

	// **参数解释**: 试用主机数量 **取值范围**: 最小值0，最大值2147483647
	TrialNum *int32 `json:"trial_num,omitempty"`

	// **参数解释**: 未防护数 **取值范围**: 最小值0，最大值2147483647
	UnprotectedNum *int32 `json:"unprotected_num,omitempty"`

	// **参数解释**: 资产风险个数 **取值范围**: 最小值0，最大值2147483647
	Asset *int32 `json:"asset,omitempty"`

	// **参数解释**: 漏洞风险个数 **取值范围**: 最小值0，最大值2147483647
	Vulnerability *int32 `json:"vulnerability,omitempty"`

	// **参数解释**: 入侵事件个数 **取值范围**: 最小值0，最大值2147483647
	Intrusion *int32 `json:"intrusion,omitempty"`

	// **参数解释**: 基线检测风险个数 **取值范围**: 最小值0，最大值2147483647
	Baseline       *int32 `json:"baseline,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowTrialsStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTrialsStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowTrialsStatisticsResponse", string(data)}, " ")
}
