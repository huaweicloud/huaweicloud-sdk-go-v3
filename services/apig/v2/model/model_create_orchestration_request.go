package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOrchestrationRequest Request Object
type CreateOrchestrationRequest struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	Body *OrchestrationCreate `json:"body,omitempty"`
}

func (o CreateOrchestrationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrchestrationRequest struct{}"
	}

	return strings.Join([]string{"CreateOrchestrationRequest", string(data)}, " ")
}
