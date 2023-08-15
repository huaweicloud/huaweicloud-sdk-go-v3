package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PhoneMigrateJob 云手机迁移任务。
type PhoneMigrateJob struct {

	// 源云手机id。
	SourcePhoneId *string `json:"source_phone_id,omitempty"`

	// 目标云手机id。
	TargetPhoneId *string `json:"target_phone_id,omitempty"`

	// 任务的唯一标识。
	JobId *string `json:"job_id,omitempty"`

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误说明。
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o PhoneMigrateJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PhoneMigrateJob struct{}"
	}

	return strings.Join([]string{"PhoneMigrateJob", string(data)}, " ")
}
