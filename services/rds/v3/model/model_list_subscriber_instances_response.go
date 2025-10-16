package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubscriberInstancesResponse Response Object
type ListSubscriberInstancesResponse struct {

	// 实例列表。
	Instances *[]ReplicationInstanceInfo `json:"instances,omitempty"`

	// 总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSubscriberInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubscriberInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListSubscriberInstancesResponse", string(data)}, " ")
}
