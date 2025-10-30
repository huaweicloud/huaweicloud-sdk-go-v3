package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEventSeverityResponse Response Object
type ShowEventSeverityResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// 低危数量
	LowNum *int32 `json:"low_num,omitempty"`

	// 中危数量
	MediumNum *int32 `json:"medium_num,omitempty"`

	// 高危数量
	HighNum *int32 `json:"high_num,omitempty"`

	// 危急数量
	CriticalNum    *int32 `json:"critical_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowEventSeverityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEventSeverityResponse struct{}"
	}

	return strings.Join([]string{"ShowEventSeverityResponse", string(data)}, " ")
}
