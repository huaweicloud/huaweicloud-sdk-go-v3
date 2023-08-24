package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AcceptOrRejectEndpointConnectionsRequest Request Object
type AcceptOrRejectEndpointConnectionsRequest struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	Body *ConnectionActionReq `json:"body,omitempty"`
}

func (o AcceptOrRejectEndpointConnectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AcceptOrRejectEndpointConnectionsRequest struct{}"
	}

	return strings.Join([]string{"AcceptOrRejectEndpointConnectionsRequest", string(data)}, " ")
}
