package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateJobConfigurationsRequest Request Object
type UpdateJobConfigurationsRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *string `json:"X-Language,omitempty"`

	Body *ModifyParameterReq `json:"body,omitempty"`
}

func (o UpdateJobConfigurationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateJobConfigurationsRequest struct{}"
	}

	return strings.Join([]string{"UpdateJobConfigurationsRequest", string(data)}, " ")
}
