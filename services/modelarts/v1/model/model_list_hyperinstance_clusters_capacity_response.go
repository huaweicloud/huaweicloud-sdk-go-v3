package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHyperinstanceClustersCapacityResponse Response Object
type ListHyperinstanceClustersCapacityResponse struct {

	// 容量信息列表
	Capacities     *[]ServerHpsClusterCapacity `json:"capacities,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListHyperinstanceClustersCapacityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHyperinstanceClustersCapacityResponse struct{}"
	}

	return strings.Join([]string{"ListHyperinstanceClustersCapacityResponse", string(data)}, " ")
}
