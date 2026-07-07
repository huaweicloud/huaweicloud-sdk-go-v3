package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartAnalysisSessionRequest Request Object
type StartAnalysisSessionRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 节点ID
	NodeId string `json:"node_id"`

	Body *StartAnalysisSessionRequestBody `json:"body,omitempty"`
}

func (o StartAnalysisSessionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartAnalysisSessionRequest struct{}"
	}

	return strings.Join([]string{"StartAnalysisSessionRequest", string(data)}, " ")
}
