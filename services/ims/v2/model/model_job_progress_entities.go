package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type JobProgressEntities struct {

	// 镜像ID
	ImageId *string `json:"image_id,omitempty"`

	// 当前任务名称
	CurrentTask *string `json:"current_task,omitempty"`

	// 镜像名称
	ImageName *string `json:"image_name,omitempty"`

	// 任务执行进度
	ProcessPercent *float64 `json:"process_percent,omitempty"`

	// 子任务ID
	SubJobId *string `json:"subJobId,omitempty"`
}

func (o JobProgressEntities) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobProgressEntities struct{}"
	}

	return strings.Join([]string{"JobProgressEntities", string(data)}, " ")
}
