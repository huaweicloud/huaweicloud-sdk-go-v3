package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StatisticsDetail struct {
	Count *CountDetail `json:"count"`

	// 起始时间
	StartTime string `json:"start_time"`
}

func (o StatisticsDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatisticsDetail struct{}"
	}

	return strings.Join([]string{"StatisticsDetail", string(data)}, " ")
}
