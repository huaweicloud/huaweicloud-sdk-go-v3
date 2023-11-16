package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Vertices struct {

	// 子任务构建记录ID
	Id *string `json:"id,omitempty"`

	// 子任务执行状态
	Status *string `json:"status,omitempty"`

	// 子任务名称
	DisplayName *string `json:"display_name,omitempty"`

	// 子任务构建耗时
	BuildDuration *int32 `json:"build_duration,omitempty"`

	// 子任务开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 子任务结束时间
	FinishTime *string `json:"finish_time,omitempty"`

	// 子任务构建编号
	BuildNo *string `json:"build_no,omitempty"`
}

func (o Vertices) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Vertices struct{}"
	}

	return strings.Join([]string{"Vertices", string(data)}, " ")
}
