package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInternalEndpointConnectionsRequest Request Object
type ListInternalEndpointConnectionsRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 返回条数，默认为10，最大值为1000。**注意：offset和limit参数需要配套使用。**
	Limit *int32 `json:"limit,omitempty"`

	// 起始索引，默认为0。**注意：offset和limit参数需要配套使用。**
	Offset *int32 `json:"offset,omitempty"`

	// 终端节点的连接状态。 - pendingAcceptance:待接受 - accepted:已接受 - rejected:已拒绝 - failed:失败
	Status *string `json:"status,omitempty"`

	// VPC终端节点ID
	Id *string `json:"id,omitempty"`
}

func (o ListInternalEndpointConnectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInternalEndpointConnectionsRequest struct{}"
	}

	return strings.Join([]string{"ListInternalEndpointConnectionsRequest", string(data)}, " ")
}
