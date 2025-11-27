package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartDbAgentJobRequest Request Object
type RestartDbAgentJobRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 作业ID
	JobId string `json:"job_id"`
}

func (o RestartDbAgentJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartDbAgentJobRequest struct{}"
	}

	return strings.Join([]string{"RestartDbAgentJobRequest", string(data)}, " ")
}
