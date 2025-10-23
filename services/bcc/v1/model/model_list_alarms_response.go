package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmsResponse Response Object
type ListAlarmsResponse struct {

	// 本次分页查询响应的告警数量
	Count *int64 `json:"count,omitempty"`

	// 下一页的marker
	NextMarker *string `json:"next_marker,omitempty"`

	// 告警数据
	Alarms         *[]AlarmEntity `json:"alarms,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListAlarmsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmsResponse struct{}"
	}

	return strings.Join([]string{"ListAlarmsResponse", string(data)}, " ")
}
