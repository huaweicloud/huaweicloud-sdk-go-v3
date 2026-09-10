package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConnectionsDetail 资源组网络连接信息列表的连接具体信息。
type ListConnectionsDetail struct {

	// 连接ID，用于标识资源组网络连接的UUID。
	Id *string `json:"id,omitempty"`

	// 创建连接时，用户自定义的连接名称。
	Name *string `json:"name,omitempty"`

	// 连接状态，包括以下三种状态： NORMAL：正常 ABNORMAL：异常 FAILED：失败
	Status *string `json:"status,omitempty"`

	// 各个集群创建对等连接的信息。
	AvailableClusterInfo *[]ConnectionClusterInfo `json:"available_cluster_info,omitempty"`

	// 对应服务的虚拟私有云标识。
	DestVpcId *string `json:"dest_vpc_id,omitempty"`

	// 对应服务的子网网络标识。
	DestNetworkId *string `json:"dest_network_id,omitempty"`

	// 创建连接的时间。为UTC的时间戳。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 用户自定义主机信息。
	Hosts *[]ConnectionsHost `json:"hosts,omitempty"`

	// 用户添加的路由信息。
	Routes *[]ConnectionsRoute `json:"routes,omitempty"`
}

func (o ListConnectionsDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConnectionsDetail struct{}"
	}

	return strings.Join([]string{"ListConnectionsDetail", string(data)}, " ")
}
