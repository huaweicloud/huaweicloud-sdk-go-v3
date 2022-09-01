package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 操作进度信息
type ProcessInfo struct {

	// 操作名
	StepName *string `json:"step_name,omitempty" xml:"step_name"`

	// 操作状态
	Status *string `json:"status,omitempty" xml:"status"`

	// 操作详情
	Detail *string `json:"detail,omitempty" xml:"detail"`

	// 子操作, \"map[string][SubDetail] key:子操作名 value:子操作结果\"
	SubSteps map[string]SubDetail `json:"sub_steps,omitempty" xml:"sub_steps"`

	// 序列号
	SerialNum *int64 `json:"serial_num,omitempty" xml:"serial_num"`
}

func (o ProcessInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProcessInfo struct{}"
	}

	return strings.Join([]string{"ProcessInfo", string(data)}, " ")
}
