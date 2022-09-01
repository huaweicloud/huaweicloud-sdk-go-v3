package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListNetworkInstancesResponse struct {

	// 网络实例列表。
	NetworkInstances *[]NetworkInstance `json:"network_instances,omitempty" xml:"network_instances"`

	PageInfo *PageInfo `json:"page_info,omitempty" xml:"page_info"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty" xml:"request_id"`
	HttpStatusCode int     `json:"-"`
}

func (o ListNetworkInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNetworkInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListNetworkInstancesResponse", string(data)}, " ")
}
