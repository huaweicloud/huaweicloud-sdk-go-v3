package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDeployedResourcesRequest Request Object
type ListDeployedResourcesRequest struct {
	Body *ListDeployedResourcesRequestBody `json:"body,omitempty"`
}

func (o ListDeployedResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDeployedResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListDeployedResourcesRequest", string(data)}, " ")
}
