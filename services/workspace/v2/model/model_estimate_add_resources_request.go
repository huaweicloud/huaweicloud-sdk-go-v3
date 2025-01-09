package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EstimateAddResourcesRequest Request Object
type EstimateAddResourcesRequest struct {
	Body *EstimateAddSubResourcesRequestBody `json:"body,omitempty"`
}

func (o EstimateAddResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EstimateAddResourcesRequest struct{}"
	}

	return strings.Join([]string{"EstimateAddResourcesRequest", string(data)}, " ")
}
