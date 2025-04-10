package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ScheduleDetailInfo struct {
}

func (o ScheduleDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScheduleDetailInfo struct{}"
	}

	return strings.Join([]string{"ScheduleDetailInfo", string(data)}, " ")
}
