package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowWorkFlowMetricRequest struct {
	// 函数工作流URN, 格式为： urn:fss:<region_id>:<project_id>:workflow:\\<package\\>:<workflow_name>:\\<version\\> 注意： package当前只支持default version当前只支持latest

	WorkflowUrn string `json:"workflow_urn"`
	// 时间段，单位为分钟

	Period *string `json:"period,omitempty"`
}

func (o ShowWorkFlowMetricRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkFlowMetricRequest struct{}"
	}

	return strings.Join([]string{"ShowWorkFlowMetricRequest", string(data)}, " ")
}
