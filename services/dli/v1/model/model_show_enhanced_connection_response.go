package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEnhancedConnectionResponse Response Object
type ShowEnhancedConnectionResponse struct {

	// 执行请求是否成功。“true”表示请求执行成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息为空。
	Message *string `json:"message,omitempty"`

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

	// 用户自定义主机信息。
	Hosts          *[]EnhancedConnectionHost `json:"hosts,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ShowEnhancedConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEnhancedConnectionResponse struct{}"
	}

	return strings.Join([]string{"ShowEnhancedConnectionResponse", string(data)}, " ")
}
