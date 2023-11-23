package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstancesSessionRequest Request Object
type ListInstancesSessionRequest struct {

	// 节点ID。
	NodeId string `json:"node_id"`

	// 索引位置，偏移量。取值大于或等于0。不传该参数时，查询偏移量默认为0，表示从最新创建的实例节点连接开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 分页查询页数。不传该参数时，默认每页50条实例节点连接信息，最大100条。
	Limit *int32 `json:"limit,omitempty"`

	// 用户端地址前缀匹配字符串。完整的地址由ip和端口号组成。不传默认查询所有。
	AddrPrefix *string `json:"addr_prefix,omitempty"`
}

func (o ListInstancesSessionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesSessionRequest struct{}"
	}

	return strings.Join([]string{"ListInstancesSessionRequest", string(data)}, " ")
}
