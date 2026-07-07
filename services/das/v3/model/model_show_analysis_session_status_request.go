package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAnalysisSessionStatusRequest Request Object
type ShowAnalysisSessionStatusRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 节点ID
	NodeId string `json:"node_id"`

	// 任务ID
	JobId string `json:"job_id"`
}

func (o ShowAnalysisSessionStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAnalysisSessionStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowAnalysisSessionStatusRequest", string(data)}, " ")
}
