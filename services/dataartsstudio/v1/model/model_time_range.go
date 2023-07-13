package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TimeRange struct {

	// 可选值：ddlUpdateTime、dataUpdateTime、ddlCreateTime
	TimeType string `json:"time_type"`

	// 开始时间
	Start string `json:"start"`

	// 结束时间
	End string `json:"end"`
}

func (o TimeRange) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TimeRange struct{}"
	}

	return strings.Join([]string{"TimeRange", string(data)}, " ")
}
