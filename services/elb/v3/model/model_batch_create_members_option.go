package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 批量添加member请求参数。
type BatchCreateMembersOption struct {
	// 后端服务器名称。

	Name *string `json:"name,omitempty"`
	// 后端云服务器的对应的IP地址，这个IP必须在subnet_cidr_id字段的子网网段中。例如：192.168.3.11。  subnet_cidr_id为空代表添加跨VPC后端，此时address必须为ipv4地址。

	Address string `json:"address"`
	// 后端服务器端口号。

	ProtocolPort int32 `json:"protocol_port"`
	// 后端云服务器所在的子网ID。该子网和后端云服务器关联的负载均衡器的子网必须在同一VPC下。  当所在LB未开启跨VPC后端，该参数必填。 当所在LB开启跨VPC后端，该参数非必填，表示添加跨VPC后端，此时address必须为ipv4地址，pool的协议必须为TCP/HTTP/HTTPS，pool所属的LB的跨VPC后端转发能力必须打开。

	SubnetCidrId *string `json:"subnet_cidr_id,omitempty"`
	// 后端云服务器的权重，请求按权重在同一后端云服务器组下的后端云服务器间分发。权重为0的后端不再接受新的请求。当后端云服务器所在的后端云服务器组的lb_algorithm的取值为SOURCE_IP时，该字段无效。

	Weight *int32 `json:"weight,omitempty"`
}

func (o BatchCreateMembersOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateMembersOption struct{}"
	}

	return strings.Join([]string{"BatchCreateMembersOption", string(data)}, " ")
}
