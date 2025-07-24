package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlarmLevelInfo struct {

	// 告警级别
	AlarmLevel *int32 `json:"alarm_level,omitempty"`

	// 告警数量
	AlarmNumber *int32 `json:"alarm_number,omitempty"`
}

func (o AlarmLevelInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmLevelInfo struct{}"
	}

	return strings.Join([]string{"AlarmLevelInfo", string(data)}, " ")
}
