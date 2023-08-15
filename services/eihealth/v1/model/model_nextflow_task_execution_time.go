package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NextflowTaskExecutionTime struct {

	// 提交时间
	Submit *string `json:"submit,omitempty"`

	// 开始时间
	Start *string `json:"start,omitempty"`

	// 完成时间
	Complete *string `json:"complete,omitempty"`

	// 总时间
	Duration *int64 `json:"duration,omitempty"`

	// 实际运行时间
	Realtime *int64 `json:"realtime,omitempty"`
}

func (o NextflowTaskExecutionTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NextflowTaskExecutionTime struct{}"
	}

	return strings.Join([]string{"NextflowTaskExecutionTime", string(data)}, " ")
}
