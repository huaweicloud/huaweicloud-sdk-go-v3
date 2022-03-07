package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAlarmResourcesRequest struct {
	// Alarm实例ID

	AlarmId string `json:"alarm_id"`
	// 偏移量

	Offset *int32 `json:"offset,omitempty"`
	// 希望的查询的数据量

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAlarmResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmResourcesRequest", string(data)}, " ")
}
