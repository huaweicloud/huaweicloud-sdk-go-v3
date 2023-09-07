package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchFlowByIdRequest Request Object
type SearchFlowByIdRequest struct {

	// flow_id
	FlowId string `json:"flow_id"`

	// 流版本
	Version *string `json:"version,omitempty"`
}

func (o SearchFlowByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchFlowByIdRequest struct{}"
	}

	return strings.Join([]string{"SearchFlowByIdRequest", string(data)}, " ")
}
