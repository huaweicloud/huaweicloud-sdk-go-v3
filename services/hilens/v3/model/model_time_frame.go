package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TimeFrame 定时任务时间段信息
type TimeFrame struct {

	// 任务开始时间
	Start string `json:"start"`

	// 任务结束时间
	Stop string `json:"stop"`
}

func (o TimeFrame) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TimeFrame struct{}"
	}

	return strings.Join([]string{"TimeFrame", string(data)}, " ")
}
