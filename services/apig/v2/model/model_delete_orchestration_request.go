package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteOrchestrationRequest Request Object
type DeleteOrchestrationRequest struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 编排规则编号
	OrchestrationId string `json:"orchestration_id"`
}

func (o DeleteOrchestrationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteOrchestrationRequest struct{}"
	}

	return strings.Join([]string{"DeleteOrchestrationRequest", string(data)}, " ")
}
