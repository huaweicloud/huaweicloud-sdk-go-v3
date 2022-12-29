package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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

	// 批量任务执行结果
	Results *[]JobEntitiesResult `json:"results,omitempty"`

	// 子任务结果列表
	SubJobsResult *[]SubJobResult `json:"sub_jobs_result,omitempty"`

	// 子任务ID列表
	SubJobsList *[]string `json:"sub_jobs_list,omitempty"`
}

func (o JobEntities) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobEntities struct{}"
	}

	return strings.Join([]string{"JobEntities", string(data)}, " ")
}
