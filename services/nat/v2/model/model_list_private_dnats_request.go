package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPrivateDnatsRequest Request Object
type ListPrivateDnatsRequest struct {

	// 功能说明：每页返回的个数。 取值范围：1~2000。 默认值：2000。
	Limit *int32 `json:"limit,omitempty"`

	// 功能说明：分页查询起始的资源ID，为空时查询第一页。 值从上一次查询的PageInfo中的next_marker或者previous_marker中获取。
	Marker *string `json:"marker,omitempty"`

	// 是否查询前一页。
	PageReverse *bool `json:"page_reverse,omitempty"`

	// DNAT规则的ID。
	Id *[]string `json:"id,omitempty"`

	// 企业项目ID。创建DNAT规则时，关联的企业项目ID。
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	// DNAT规则的描述。长度范围小于等于255个字符，不能包含“<”和“>”。
	Description *[]string `json:"description,omitempty"`

	// 私网NAT网关实例的ID。
	GatewayId *[]string `json:"gateway_id,omitempty"`

	// 中转IP的ID。
	TransitIpId *[]string `json:"transit_ip_id,omitempty"`

	// 计算实例、ELBV2、ELBV3、VIP等资源的端口ID。
	NetworkInterfaceId *[]string `json:"network_interface_id,omitempty"`

	// DNAT规则后端的类型。 取值：     COMPUTE：后端为计算实例。     VIP：后端为VIP的实例。     ELB：后端为ELBv2的实例。     ELBv3：后端为ELBv3的实例。     CUSTOMIZE：后端为自定义IP。
	Type *[]string `json:"type,omitempty"`

	// 后端资源（计算实例、ELBV2、ELBV3、VIP等）的私网IP地址。
	PrivateIpAddress *[]string `json:"private_ip_address,omitempty"`

	// DNAT规则协议类型， 目前支持TCP/tcp/Tcp/tCp/tcP/TCp/tCP/TcP、 UDP/udp/Udp/uDp/udP/UDp/uDP/UdP、 ANY/any/Any/aNy/anY/ANy/aNY/AnY。 分别对应协议号6、17、0。
	Protocol *[]string `json:"protocol,omitempty"`

	// 后端实例的端口号（计算实例、ELBV2、ELBV3、VIP等)。
	InternalServicePort *[]string `json:"internal_service_port,omitempty"`

	// 中转IP的端口号。
	TransitServicePort *[]string `json:"transit_service_port,omitempty"`

	// 中转IP的地址。
	TransitIpAddress *[]string `json:"transit_ip_address,omitempty"`
}

func (o ListPrivateDnatsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivateDnatsRequest struct{}"
	}

	return strings.Join([]string{"ListPrivateDnatsRequest", string(data)}, " ")
}
