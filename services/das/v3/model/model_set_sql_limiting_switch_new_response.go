package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetSqlLimitingSwitchNewResponse Response Object
type SetSqlLimitingSwitchNewResponse struct {

	// 开关状态
	SwitchOn *string `json:"switch_on,omitempty"`

	// 是否需要重试
	Retry *bool `json:"retry,omitempty"`

	// 错误信息
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 状态
	Status *bool `json:"status,omitempty"`

	// 详细状态
	DetailStatus *string `json:"detail_status,omitempty"`

	// 失败原因
	FailReason *string `json:"fail_reason,omitempty"`

	// 工作流ID
	JobId *string `json:"job_id,omitempty"`

	// 工作流状态
	JobStatus      *string `json:"job_status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetSqlLimitingSwitchNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetSqlLimitingSwitchNewResponse struct{}"
	}

	return strings.Join([]string{"SetSqlLimitingSwitchNewResponse", string(data)}, " ")
}
