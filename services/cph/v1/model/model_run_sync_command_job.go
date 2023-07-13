package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunSyncCommandJob 执行异步shell命令任务。
type RunSyncCommandJob struct {

	// 云手机的唯一标识，云手机相关任务包含此字段
	PhoneId *string `json:"phone_id,omitempty"`

	// 任务的唯一标识
	JobId *string `json:"job_id,omitempty"`

	// 任务状态 - 2：成功 - 1：运行中 - -1：失败
	Status *int32 `json:"status,omitempty"`

	// 任务错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 任务错误码说明
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 任务执行返回内容，最长1024字节
	ExecuteMsg *string `json:"execute_msg,omitempty"`
}

func (o RunSyncCommandJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunSyncCommandJob struct{}"
	}

	return strings.Join([]string{"RunSyncCommandJob", string(data)}, " ")
}
