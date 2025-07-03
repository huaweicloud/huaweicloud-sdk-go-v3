package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyFlowSourcesRequest Request Object
type ModifyFlowSourcesRequest struct {

	// 流id
	FlowId string `json:"flow_id"`

	// 流源名称
	SourceName string `json:"source_name"`

	Body *ModifyFlowSourcesRequestBody `json:"body,omitempty"`
}

func (o ModifyFlowSourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyFlowSourcesRequest struct{}"
	}

	return strings.Join([]string{"ModifyFlowSourcesRequest", string(data)}, " ")
}
