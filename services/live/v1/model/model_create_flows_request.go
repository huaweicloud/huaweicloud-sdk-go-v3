package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFlowsRequest Request Object
type CreateFlowsRequest struct {
	Body *CreateFlowsRequestBody `json:"body,omitempty"`
}

func (o CreateFlowsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlowsRequest struct{}"
	}

	return strings.Join([]string{"CreateFlowsRequest", string(data)}, " ")
}
