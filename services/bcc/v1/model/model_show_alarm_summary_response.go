package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAlarmSummaryResponse Response Object
type ShowAlarmSummaryResponse struct {

	// 告警总数
	Count *int64 `json:"count,omitempty"`

	// 告警统计信息
	Statistics     *[]ShowAlarmSummaryResponseBodyStatistics `json:"statistics,omitempty"`
	HttpStatusCode int                                       `json:"-"`
}

func (o ShowAlarmSummaryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlarmSummaryResponse struct{}"
	}

	return strings.Join([]string{"ShowAlarmSummaryResponse", string(data)}, " ")
}
