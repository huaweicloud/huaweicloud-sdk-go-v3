package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 操作进度信息
type ProcessInfo struct {
	// 操作名

	StepName *string `json:"step_name,omitempty"`
	// 操作状态

	Status *string `json:"status,omitempty"`
	// 操作详情

	Detail *string `json:"detail,omitempty"`
	// 子操作, \"map[string][SubDetail] key:子操作名 value:子操作结果\"

	SubSteps map[string]SubDetail `json:"sub_steps,omitempty"`
	// 序列号

	SerialNum *int64 `json:"serial_num,omitempty"`
}

func (o ProcessInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProcessInfo struct{}"
	}

	return strings.Join([]string{"ProcessInfo", string(data)}, " ")
}
