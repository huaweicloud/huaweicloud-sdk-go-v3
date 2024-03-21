package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StatisticsPvo struct {

	// 结束时间。
	EndTime string `json:"endTime"`

	// 开始时间。
	StartTime string `json:"startTime"`
}

func (o StatisticsPvo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatisticsPvo struct{}"
	}

	return strings.Join([]string{"StatisticsPvo", string(data)}, " ")
}
