package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSqlLimitingJobInfoResponse Response Object
type ShowSqlLimitingJobInfoResponse struct {

	// 任务ID
	JobId *string `json:"job_id,omitempty"`

	// 任务名称
	JobName *string `json:"job_name,omitempty"`

	// 任务类型
	JobType *string `json:"job_type,omitempty"`

	// 任务状态
	Status *string `json:"status,omitempty"`

	// 失败原因
	FailReason     *string `json:"fail_reason,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSqlLimitingJobInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlLimitingJobInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowSqlLimitingJobInfoResponse", string(data)}, " ")
}
