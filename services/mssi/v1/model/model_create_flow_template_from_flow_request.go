package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFlowTemplateFromFlowRequest Request Object
type CreateFlowTemplateFromFlowRequest struct {

	// ID of flow
	FlowId string `json:"flow_id"`

	Body *TemplateMessage `json:"body,omitempty"`
}

func (o CreateFlowTemplateFromFlowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlowTemplateFromFlowRequest struct{}"
	}

	return strings.Join([]string{"CreateFlowTemplateFromFlowRequest", string(data)}, " ")
}
