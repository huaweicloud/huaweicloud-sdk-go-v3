package model

import (
	"encoding/json"

	"strings"
)

//
type JobEntities struct {
	// 镜像ID

	ImageId *string `json:"image_id,omitempty"`
	// 当前任务名称

	CurrentTask *string `json:"current_task,omitempty"`
	// 镜像名称

	ImageName *string `json:"image_name,omitempty"`
	// 任务执行进度

	ProcessPercent *float64 `json:"process_percent,omitempty"`
}

func (o JobEntities) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "JobEntities struct{}"
	}

	return strings.Join([]string{"JobEntities", string(data)}, " ")
}
