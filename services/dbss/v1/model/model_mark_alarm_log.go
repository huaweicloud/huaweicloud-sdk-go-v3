package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MarkAlarmLog struct {

	// 告警ID列表
	Ids []string `json:"ids"`
}

func (o MarkAlarmLog) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MarkAlarmLog struct{}"
	}

	return strings.Join([]string{"MarkAlarmLog", string(data)}, " ")
}
