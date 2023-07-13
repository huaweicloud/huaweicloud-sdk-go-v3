package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DtSchedule 调度策略
type DtSchedule struct {

	// 调度周期，正则： \"1m|5m|15m|1h\"，表示从每小时第0分0秒开始，每1m、5m、15m、1h调度
	Period string `json:"period"`
}

func (o DtSchedule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DtSchedule struct{}"
	}

	return strings.Join([]string{"DtSchedule", string(data)}, " ")
}
