package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 后端云服务器响应体
type MemberResp struct {
	// 后端云服务器ID

	Id string `json:"id"`
	// 后端云服务器所在的项目ID。

	ProjectId string `json:"project_id"`
	// 后端云服务器所在的项目ID。

	TenantId string `json:"tenant_id"`
	// 后端云服务器名称。

	Name string `json:"name"`
	// 后端云服务器的管理状态；该字段虽然支持创建、更新，但实际取值决定于后端云服务器对应的弹性云服务器是否存在。该字段虽然支持创建、更新，但实际取值决定于member对应的弹性云服务器是否存在。若存在，该值为true，否则，该值为false。

	AdminStateUp bool `json:"admin_state_up"`
	// 后端端口和协议号

	ProtocolPort int32 `json:"protocol_port"`
	// 后端云服务器所在的子网ID。该子网和后端云服务器关联的负载均衡器的子网必须在同一VPC下。只支持指定IPv4的子网ID。暂不支持IPv6。

	SubnetId string `json:"subnet_id"`
	// 后端云服务器的对应的IP地址，这个IP必须在subnet_id字段的子网网段中。例如：192.168.3.11。只能指定为主网卡的IP。

	Address string `json:"address"`
	// 后端云服务器的权重，请求按权重在同一后端云服务器组下的后端云服务器间分发。权重为0的后端不再接受新的请求。当后端云服务器所在的后端云服务器组的lb_algorithm的取值为SOURCE_IP时，该字段无效。

	Weight int32 `json:"weight"`
	// 后端云服务器的健康状态，取值：  ONLINE：健康检查在线，后端服务正常。 OFFLINE：健康检查离线，后端服务异常，负载均衡器不再向异常的后端发送流量。 NO_MONITOR：无检查检查。未创建检查检查或健康检查的admin_state_up字段为false。

	OperatingStatus string `json:"operating_status"`
}

func (o MemberResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MemberResp struct{}"
	}

	return strings.Join([]string{"MemberResp", string(data)}, " ")
}
