package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDatasourceConnectionResponse Response Object
type ShowDatasourceConnectionResponse struct {

	// 执行请求是否成功。“true”表示请求执行成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息为空。
	Message *string `json:"message,omitempty"`

	// 连接ID，用于标识资源组网络连接的UUID。
	Id *string `json:"id,omitempty"`

	// 创建连接时，用户自定义的连接名称。
	Name *string `json:"name,omitempty"`

	// 连接状态，包括以下两种状态： ACTIVE：已激活 DELETED：已删除
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
	Routes         *[]ConnectionsRoute `json:"routes,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowDatasourceConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDatasourceConnectionResponse struct{}"
	}

	return strings.Join([]string{"ShowDatasourceConnectionResponse", string(data)}, " ")
}
