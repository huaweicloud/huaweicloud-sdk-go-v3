package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDistributorInstancesResponse Response Object
type ListDistributorInstancesResponse struct {

	// 实例列表。
	Instances *[]ReplicationInstanceInfo `json:"instances,omitempty"`

	// 总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDistributorInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDistributorInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListDistributorInstancesResponse", string(data)}, " ")
}
