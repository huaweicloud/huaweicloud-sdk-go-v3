package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyDbAgentJobRequest Request Object
type ModifyDbAgentJobRequest struct {

	// 语言。默认en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 作业ID
	JobId string `json:"job_id"`

	Body *ModifyDbAgentJobRequestBody `json:"body,omitempty"`
}

func (o ModifyDbAgentJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyDbAgentJobRequest struct{}"
	}

	return strings.Join([]string{"ModifyDbAgentJobRequest", string(data)}, " ")
}
