package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 子操作详情
type SubDetail struct {

	// 子操作名
	SubstepName *string `json:"substep_name,omitempty"`

	// 子操作详情
	Detail *string `json:"detail,omitempty"`

	// 子操作状态
	Status *string `json:"status,omitempty"`

	// 子操作过程信息记录
	Message *[]string `json:"message,omitempty"`

	// 开始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 结束时间
	EndTime *int64 `json:"end_time,omitempty"`

	// 序列号
	SerialNum *int64 `json:"serial_num,omitempty"`
}

func (o SubDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubDetail struct{}"
	}

	return strings.Join([]string{"SubDetail", string(data)}, " ")
}
