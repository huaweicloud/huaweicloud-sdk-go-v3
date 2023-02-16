package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateDiagnoseJobResponse struct {

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 工具任务ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDiagnoseJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDiagnoseJobResponse struct{}"
	}

	return strings.Join([]string{"CreateDiagnoseJobResponse", string(data)}, " ")
}
