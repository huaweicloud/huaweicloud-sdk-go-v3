package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Job 任务。
type Job struct {

	// 云手机的唯一标识，云手机相关任务包含此字段
	PhoneId *string `json:"phone_id,omitempty"`

	// 云手机服务器的唯一标识ID，云手机服务器相关任务包含此字段
	ServerId *string `json:"server_id,omitempty"`

	// （已废弃）云手机服务器的唯一标识ID，云手机服务 含此字段
	NodeId *string `json:"node_id,omitempty"`

	// 任务的唯一标识
	JobId *string `json:"job_id,omitempty"`

	// 任务处理开始时间 时间格式为UTC，YYYY-MM-DDTHH:MM:SSZ
	BeginTime *string `json:"begin_time,omitempty"`

	// 任务处理结束时间 时间格式为UTC，YYYY-MM-DDTHH:MM:SSZ
	EndTime *string `json:"end_time,omitempty"`

	// 任务状态 - 1 ：运行中 - 2 ： 成功 - -1 ：失败
	Status *int32 `json:"status,omitempty"`

	// 任务错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 任务错误码说明
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 任务执行返回内容，最长1024个字节
	ExecuteMsg *string `json:"execute_msg,omitempty"`
}

func (o Job) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Job struct{}"
	}

	return strings.Join([]string{"Job", string(data)}, " ")
}
