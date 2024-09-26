package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNetworkInstancesResponse Response Object
type ListNetworkInstancesResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 网络实例列表。
	NetworkInstances []NetworkInstance `json:"network_instances"`
	HttpStatusCode   int               `json:"-"`
}

func (o ListNetworkInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNetworkInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListNetworkInstancesResponse", string(data)}, " ")
}
