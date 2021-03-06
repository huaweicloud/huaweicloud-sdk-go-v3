package model

import (
	"encoding/json"

	"strings"
)

// 后端云服务器组列表查询返回对象。
type Member struct {
	// 后端云服务器ID。

	Id string `json:"id"`
	// 后端云服务器名称。

	Name string `json:"name"`
	// 后端云服务器所在的项目ID。

	ProjectId string `json:"project_id"`
	// 所属服务器组ID。  注意：该字段当前仅GET /v3/{project_id}/elb/members 接口可见。

	PoolId *string `json:"pool_id,omitempty"`
	// 后端云服务器的管理状态；该字段虽然支持创建、更新，但实际取值决定于后端云服务器对应的弹性云服务器是否存在。若存在，该值为true，否则，该值为false。

	AdminStateUp bool `json:"admin_state_up"`
	// 后端云服务器所在的子网ID。该子网和后端云服务器关联的负载均衡器的子网必须在同一VPC下。只支持指定IPv4的子网ID。暂不支持IPv6。 为空代表当前member为跨VPC后端

	SubnetCidrId *string `json:"subnet_cidr_id,omitempty"`
	// 后端服务器端口号

	ProtocolPort int32 `json:"protocol_port"`
	// 后端云服务器的权重，请求按权重在同一后端云服务器组下的后端云服务器间分发。权重为0的后端不再接受新的请求。当后端云服务器所在的后端云服务器组的lb_algorithm的取值为SOURCE_IP时，该字段无效。

	Weight int32 `json:"weight"`
	// 后端云服务器的对应的IP地址，这个IP必须在subnet_cidr_id字段的子网网段中。例如：192.168.3.11。只能指定为主网卡的IP。

	Address string `json:"address"`
	// 只读属性，根据传入的address字段自动判断之后生成，取值范围(v4、v6)

	IpVersion string `json:"ip_version"`
	// 设备使用者，为空表示后端服务器未关联到ECS。  注意：该字段当前仅GET /v3/{project_id}/elb/members 接口可见。

	DeviceOwner *string `json:"device_owner,omitempty"`
	// 关联的ECS ID，为空表示后端服务器未关联到ECS。  注意：该字段当前仅GET /v3/{project_id}/elb/members 接口可见。

	DeviceId *string `json:"device_id,omitempty"`
	// 后端云服务器的健康状态，可以为ONLINE，NO_MONITOR，OFFLINE。

	OperatingStatus string `json:"operating_status"`
	// 所属负载均衡器ID。  注意：该字段当前仅GET /v3/{project_id}/elb/members 接口可见。

	LoadbalancerId *string `json:"loadbalancer_id,omitempty"`
}

func (o Member) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Member struct{}"
	}

	return strings.Join([]string{"Member", string(data)}, " ")
}
