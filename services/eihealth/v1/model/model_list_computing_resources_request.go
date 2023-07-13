package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListComputingResourcesRequest Request Object
type ListComputingResourcesRequest struct {
}

func (o ListComputingResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListComputingResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListComputingResourcesRequest", string(data)}, " ")
}
