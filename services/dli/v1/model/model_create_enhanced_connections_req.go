package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEnhancedConnectionsReq 创建增强型跨源连接的请求body体。
type CreateEnhancedConnectionsReq struct {

	// 连接名称。长度64，数字字母下划线组成。
	Name string `json:"name"`

	// 对应服务的vpc的ID。
	DestVpcId string `json:"dest_vpc_id"`

	// 对应服务的子网网络ID，即为需要建立连接的服务所在的子网。
	DestNetworkId string `json:"dest_network_id"`

	// 需要使用跨源的队列列表。
	Queues *[]string `json:"queues,omitempty"`

	// 用户自定义主机信息，最大支持2万条记录。
	Hosts *[]EnhancedConnectionsHost `json:"hosts,omitempty"`

	// 对应服务的子网关联的路由表。
	RoutetableId *string `json:"routetable_id,omitempty"`

	// 标签
	Tags *[]TmsTagEntity `json:"tags,omitempty"`
}

func (o CreateEnhancedConnectionsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEnhancedConnectionsReq struct{}"
	}

	return strings.Join([]string{"CreateEnhancedConnectionsReq", string(data)}, " ")
}
