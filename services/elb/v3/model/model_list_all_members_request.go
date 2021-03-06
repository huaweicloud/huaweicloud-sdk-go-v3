package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListAllMembersRequest struct {
	// 后端云服务器的对应的IP地址，这个IP必须在subnet_cidr_id字段的子网网段中，例如：192.168.3.11。只能指定为主网卡的IP。

	Address *[]string `json:"address,omitempty"`
	// 后端云服务器的管理状态。该字段虽然支持创建、更新，但实际取值决定于member对应的弹性云服务器是否存在。若存在，该值为true，否则，该值为false。

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// 企业项目ID。

	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`
	// 后端云服务器ID。

	Id *[]string `json:"id,omitempty"`
	// IP版本，取值v4、v6

	IpVersion *string `json:"ip_version,omitempty"`
	// 每页返回的个数。

	Limit *int32 `json:"limit,omitempty"`
	// member所属的负载均衡器ID

	LoadbalancerId *string `json:"loadbalancer_id,omitempty"`
	// 上一页最后一条记录的ID。  使用说明：  - 必须与limit一起使用。 - 不指定时表示查询第一页。 - 该字段不允许为空或无效的ID。

	Marker *string `json:"marker,omitempty"`
	// 后端云服务器名称。

	Name *[]string `json:"name,omitempty"`
	// 后端云服务器的健康状态，可以为ONLINE，NO_MONITOR，OFFLINE。

	OperatingStatus *[]string `json:"operating_status,omitempty"`
	// 分页的顺序，true表示从后往前分页，false表示从前往后分页，默认为false。  使用说明：必须与limit一起使用。

	PageReverse *bool `json:"page_reverse,omitempty"`
	// member所属的服务器组ID

	PoolId *string `json:"pool_id,omitempty"`
	// 后端端口和协议号。

	ProtocolPort *[]int32 `json:"protocol_port,omitempty"`
	// 后端云服务器所在的子网ID。该子网和后端云服务器关联的负载均衡器的子网必须在同一VPC下。只支持指定IPv4的子网ID。暂不支持IPv6。

	SubnetCidrId *[]string `json:"subnet_cidr_id,omitempty"`
	// 后端云服务器的权重，请求按权重在同一后端云服务器组下的后端云服务器间分发。权重为0的后端不再接受新的请求。当后端云服务器所在的后端云服务器组的lb_algorithm的取值为SOURCE_IP时，该字段无效。

	Weight *[]int32 `json:"weight,omitempty"`
}

func (o ListAllMembersRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListAllMembersRequest struct{}"
	}

	return strings.Join([]string{"ListAllMembersRequest", string(data)}, " ")
}
