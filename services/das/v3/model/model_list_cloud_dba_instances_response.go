package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudDbaInstancesResponse Response Object
type ListCloudDbaInstancesResponse struct {

	// 实例列表。
	InstanceList *[]DasInstanceInfo `json:"instance_list,omitempty"`

	// 总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListCloudDbaInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudDbaInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListCloudDbaInstancesResponse", string(data)}, " ")
}
