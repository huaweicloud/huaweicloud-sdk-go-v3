package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFactoryAlarmInfoResponse Response Object
type ListFactoryAlarmInfoResponse struct {

	// 通知记录信息
	AlarmInfo *[]AlarmInfoResponseAlarmInfo `json:"alarm_info,omitempty"`

	// 通知记录数量
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListFactoryAlarmInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFactoryAlarmInfoResponse struct{}"
	}

	return strings.Join([]string{"ListFactoryAlarmInfoResponse", string(data)}, " ")
}
