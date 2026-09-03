package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExecutionTime struct {

	// 开始时间 00:00
	BeginTime *string `json:"beginTime,omitempty"`

	// 结束时间 23:59
	EndTime *string `json:"endTime,omitempty"`
}

func (o ExecutionTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecutionTime struct{}"
	}

	return strings.Join([]string{"ExecutionTime", string(data)}, " ")
}
