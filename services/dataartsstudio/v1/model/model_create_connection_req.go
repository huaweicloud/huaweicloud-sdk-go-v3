package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConnectionReq 创建资源组网络连接的请求body体。
type CreateConnectionReq struct {

	// 连接名称。长度64，数字字母下划线组成。
	Name string `json:"name"`

	// 对应服务的vpc的ID。
	DestVpcId string `json:"dest_vpc_id"`

	// 对应服务的子网网络ID，即为需要建立连接的服务所在的子网。
	DestNetworkId string `json:"dest_network_id"`

	// 需要使用连接的集群ID列表。单条最大长度128字符。
	Clusters *[]string `json:"clusters,omitempty"`

	// 用户自定义主机信息，最大支持2万条记录。
	Hosts *[]ConnectionsHost `json:"hosts,omitempty"`

	// 对应服务的子网关联的路由表。
	RoutetableId *string `json:"routetable_id,omitempty"`
}

func (o CreateConnectionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConnectionReq struct{}"
	}

	return strings.Join([]string{"CreateConnectionReq", string(data)}, " ")
}
