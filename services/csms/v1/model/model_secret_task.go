package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SecretTask struct {

	// 任务ID
	TaskId *string `json:"task_id,omitempty"`

	// 凭据名称。
	SecretName *string `json:"secret_name,omitempty"`

	// FunctionGraph函数的urn。
	RotationFuncUrn *string `json:"rotation_func_urn,omitempty"`

	// 任务状态。
	TaskStatus *string `json:"task_status,omitempty"`

	// 任务尝试次数。
	AttemptNums *int32 `json:"attempt_nums,omitempty"`

	// 轮转类型。
	OperateType *string `json:"operate_type,omitempty"`

	// 任务创建时间。
	TaskTime *int64 `json:"task_time,omitempty"`

	// 任务错误码。
	TaskErrorCode *string `json:"task_error_code,omitempty"`

	// 任务错误信息。
	TaskErrorMsg *string `json:"task_error_msg,omitempty"`
}

func (o SecretTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecretTask struct{}"
	}

	return strings.Join([]string{"SecretTask", string(data)}, " ")
}
