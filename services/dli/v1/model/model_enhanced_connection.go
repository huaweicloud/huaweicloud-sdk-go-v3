package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnhancedConnection 跨源连接信息列表的连接具体信息。
type EnhancedConnection struct {

	// 连接ID，用于标识跨源连接的UUID。
	Id *string `json:"id,omitempty"`

	// 创建连接时，用户自定义的连接名称。
	Name *string `json:"name,omitempty"`

	// 连接状态，包括以下两种状态： ACTIVE：已激活 DELETED：已删除
	Status *string `json:"status,omitempty"`

	// 各个队列创建对等连接的信息。
	AvailableQueueInfo *[]EnhancedConnectionResource `json:"available_queue_info,omitempty"`

	// 各个弹性资源池创建对等连接的信息。
	ElasticResourcePools *[]EnhancedConnectionResource `json:"elastic_resource_pools,omitempty"`

	// 对应服务的虚拟私有云标识。
	DestVpcId *string `json:"dest_vpc_id,omitempty"`

	// 对应服务的子网网络标识。
	DestNetworkId *string `json:"dest_network_id,omitempty"`

	// 创建连接的时间。为UTC的时间戳。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 用户自定义主机信息
	Hosts *[]EnhancedConnectionHost `json:"hosts,omitempty"`

	// 该增强跨源连接如果做过项目赋权，则该字段是\"false\"，否则为\"true\"。
	IsPrivis *bool `json:"isPrivis,omitempty"`
}

func (o EnhancedConnection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnhancedConnection struct{}"
	}

	return strings.Join([]string{"EnhancedConnection", string(data)}, " ")
}
