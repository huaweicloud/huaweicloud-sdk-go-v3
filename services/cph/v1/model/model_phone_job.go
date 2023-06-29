package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PhoneJob 云手机任务。
type PhoneJob struct {

	// 云手机的唯一标识ID，云手机相关任务包含此字段。
	PhoneId *string `json:"phone_id,omitempty"`

	// 任务的唯一标识。
	JobId *string `json:"job_id,omitempty"`

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误说明。
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o PhoneJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PhoneJob struct{}"
	}

	return strings.Join([]string{"PhoneJob", string(data)}, " ")
}
