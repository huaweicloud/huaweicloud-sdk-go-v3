package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskRequest 创建或更新作业请求体
type TaskRequest struct {

	// 作业动作，创建作业或者是删除作业又或是更新作业等
	Business *string `json:"business,omitempty"`

	Data *TaskData `json:"data"`

	// 作业描述
	Description *string `json:"description,omitempty"`

	// 实例ID，非必选
	InstanceId *string `json:"instance_id,omitempty"`

	// 作业名称
	Name string `json:"name"`

	// 时间戳，非必选
	Timestamp *string `json:"timestamp,omitempty"`
}

func (o TaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskRequest struct{}"
	}

	return strings.Join([]string{"TaskRequest", string(data)}, " ")
}
