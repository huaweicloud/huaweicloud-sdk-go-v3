package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchUpdateConfigsResponse struct {
	// 批量任务ID，创建批量任务时由物联网平台分配获得。

	TaskId *string `json:"task_id,omitempty"`
	// 批量任务名称。

	TaskName *string `json:"task_name,omitempty"`
	// 任务类型。

	TaskType *string `json:"task_type,omitempty"`
	// 批量任务的状态，可选参数，取值范围：Success|Fail|Processing|PartialSuccess|Stopped|Waitting|Initializing。

	Status *string `json:"status,omitempty"`
	// 批量任务状态描述(包含主任务失败错误信息)

	StatusDesc     *string `json:"status_desc,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchUpdateConfigsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchUpdateConfigsResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateConfigsResponse", string(data)}, " ")
}
