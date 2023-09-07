package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFlowRequest Request Object
type CreateFlowRequest struct {
	Body *FlowMeta `json:"body,omitempty"`
}

func (o CreateFlowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlowRequest struct{}"
	}

	return strings.Join([]string{"CreateFlowRequest", string(data)}, " ")
}
