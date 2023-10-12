package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OneClickAlarmId 一键告警ID
type OneClickAlarmId struct {
}

func (o OneClickAlarmId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OneClickAlarmId struct{}"
	}

	return strings.Join([]string{"OneClickAlarmId", string(data)}, " ")
}
