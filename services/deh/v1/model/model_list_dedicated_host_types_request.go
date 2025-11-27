package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDedicatedHostTypesRequest Request Object
type ListDedicatedHostTypesRequest struct {

	// AZ。
	AvailabilityZone string `json:"availability_zone"`

	// 查询返回云服务器列表当前页面的数量。
	Limit *string `json:"limit,omitempty"`

	// 以单页最后一条专属主机的host_type作为分页标记
	Marker *string `json:"marker,omitempty"`
}

func (o ListDedicatedHostTypesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDedicatedHostTypesRequest struct{}"
	}

	return strings.Join([]string{"ListDedicatedHostTypesRequest", string(data)}, " ")
}
