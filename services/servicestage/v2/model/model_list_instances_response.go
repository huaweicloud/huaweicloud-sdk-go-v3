package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListInstancesResponse struct {
	// 实例总数。

	Count *int32 `json:"count,omitempty"`
	// 实例列表。

	Instances      *[]InstanceListView `json:"instances,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListInstancesResponse", string(data)}, " ")
}
