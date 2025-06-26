package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskDetail struct {

	// 任务ID
	Id int32 `json:"id"`

	// 执行任务的ID
	ExecutionId int32 `json:"execution_id"`

	// 资源类型
	ResourceType string `json:"resource_type"`

	// 源资源
	SrcResource string `json:"src_resource"`

	// 目标资源
	DstResource string `json:"dst_resource"`

	// 操作类型
	Operation string `json:"operation"`

	// harbor任务ID
	JobId string `json:"job_id"`

	// 状态
	Status string `json:"status"`

	// 状态修订
	StatusRevision int32 `json:"StatusRevision"`

	// 开始时间
	StartTime string `json:"start_time"`

	// 结束时间
	EndTime string `json:"end_time"`
}

func (o TaskDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskDetail struct{}"
	}

	return strings.Join([]string{"TaskDetail", string(data)}, " ")
}
