package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateMembersOption 批量添加member请求参数。
type BatchCreateMembersOption struct {

	// 后端服务器名称。注意：该名称并非ECS名称。
	Name *string `json:"name,omitempty"`

	// 后端云服务器的对应的IP地址，这个IP必须在subnet_cidr_id字段的子网网段中。例如：192.168.3.11。  subnet_cidr_id为空代表添加跨VPC后端，此时address必须为ipv4地址。
	Address string `json:"address"`

	// 后端服务器端口号。
	ProtocolPort int32 `json:"protocol_port"`

	// 后端云服务器所在的子网ID。该子网和后端云服务器关联的负载均衡器的子网必须在同一VPC下。  当所在LB未开启跨VPC后端，该参数必填。  当所在LB开启跨VPC后端，该参数非必填，表示添加跨VPC后端，此时address必须为ipv4地址， pool的协议必须为TCP/HTTP/HTTPS，pool所属的LB的跨VPC后端转发能力必须打开。
	SubnetCidrId *string `json:"subnet_cidr_id,omitempty"`

	// 后端云服务器的权重，请求将根据pool配置的负载均衡算法和后端云服务器的权重进行负载分发。 权重值越大，分发的请求越多。权重为0的后端不再接受新的请求。  取值：0-100，默认1。  使用说明：若所在pool的lb_algorithm取值为SOURCE_IP，该字段无效。
	Weight *int32 `json:"weight,omitempty"`
}

func (o BatchCreateMembersOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateMembersOption struct{}"
	}

	return strings.Join([]string{"BatchCreateMembersOption", string(data)}, " ")
}
