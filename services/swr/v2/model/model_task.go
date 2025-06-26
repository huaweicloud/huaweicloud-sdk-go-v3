package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Task struct {

	// 老化策略执行记录任务ID
	Id int32 `json:"id"`

	// 老化策略执行记录ID
	ExecutionId int32 `json:"execution_id"`

	// 制品仓库名称
	Repository string `json:"repository"`

	// 任务名称
	JobId string `json:"job_id"`

	// 执行状态
	Status string `json:"status"`

	// 状态码
	StatusCode int32 `json:"status_code"`

	// 状态修订
	StatusRevision int32 `json:"status_revision"`

	// 开始时间
	StartTime string `json:"start_time"`

	// 结束时间
	EndTime string `json:"end_time"`

	// 版本总数
	Total int32 `json:"total"`

	// 保留版本总数
	Retained int32 `json:"retained"`
}

func (o Task) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Task struct{}"
	}

	return strings.Join([]string{"Task", string(data)}, " ")
}
