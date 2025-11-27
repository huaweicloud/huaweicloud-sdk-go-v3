package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchDbAgentJobRequest Request Object
type SwitchDbAgentJobRequest struct {

	// 语言。默认en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 作业ID
	JobId string `json:"job_id"`
}

func (o SwitchDbAgentJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchDbAgentJobRequest struct{}"
	}

	return strings.Join([]string{"SwitchDbAgentJobRequest", string(data)}, " ")
}
