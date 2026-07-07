package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAnalysisSessionResultRequest Request Object
type ShowAnalysisSessionResultRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 节点ID
	NodeId string `json:"node_id"`

	// 任务ID
	JobId string `json:"job_id"`
}

func (o ShowAnalysisSessionResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAnalysisSessionResultRequest struct{}"
	}

	return strings.Join([]string{"ShowAnalysisSessionResultRequest", string(data)}, " ")
}
