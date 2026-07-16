package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHyperinstanceClustersCapacityRequest Request Object
type ListHyperinstanceClustersCapacityRequest struct {
	Body *HyperinstanceClustersCapacityRequest `json:"body,omitempty"`
}

func (o ListHyperinstanceClustersCapacityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHyperinstanceClustersCapacityRequest struct{}"
	}

	return strings.Join([]string{"ListHyperinstanceClustersCapacityRequest", string(data)}, " ")
}
