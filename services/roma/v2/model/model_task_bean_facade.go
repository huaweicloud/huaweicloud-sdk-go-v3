package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskBeanFacade 任务执行详情
type TaskBeanFacade struct {

	// 任务ID
	TaskId *string `json:"task_id,omitempty"`

	// 任务名称
	TaskName *string `json:"task_name,omitempty"`

	// 失败的错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误详情
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o TaskBeanFacade) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskBeanFacade struct{}"
	}

	return strings.Join([]string{"TaskBeanFacade", string(data)}, " ")
}
