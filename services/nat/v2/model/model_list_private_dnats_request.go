package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPrivateDnatsRequest Request Object
type ListPrivateDnatsRequest struct {

	// 功能说明：每页返回的个数。 取值范围：0~2000。 默认值：2000。
	Limit *int32 `json:"limit,omitempty"`

	// 功能说明：分页查询起始的资源ID，为空时查询第一页。 值从上一次查询的PageInfo中的next_marker或者previous_marker中获取。
	Marker *string `json:"marker,omitempty"`

	// 是否查询前一页。
	PageReverse *bool `json:"page_reverse,omitempty"`

	// DNAT规则的ID。
	Id *[]string `json:"id,omitempty"`

	// 企业项目ID。创建DNAT规则时，关联的企业项目ID。
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	// DNAT规则的描述。
	Description *[]string `json:"description,omitempty"`

	// 私网NAT网关实例的ID。
	GatewayId *[]string `json:"gateway_id,omitempty"`

	// 中转IP的ID。
	TransitIpId *[]string `json:"transit_ip_id,omitempty"`

	// 中转IP的地址。
	ExternalIpAddress *[]string `json:"external_ip_address,omitempty"`

	// 网络接口ID，支持计算、ELB、VIP等实例的网络接口。
	NetworkInterfaceId *[]string `json:"network_interface_id,omitempty"`

	// DNAT规则后端的类型。 取值：     COMPUTE：后端为计算实例。     VIP：后端为VIP的实例。     ELB：后端为ELB的实例。     ELBv3：后端为ELBv3的实例。     CUSTOMIZE：后端为自定义IP。
	Type *[]string `json:"type,omitempty"`

	// 后端实例的IP私网地址。
	PrivateIpAddress *[]string `json:"private_ip_address,omitempty"`
}

func (o ListPrivateDnatsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivateDnatsRequest struct{}"
	}

	return strings.Join([]string{"ListPrivateDnatsRequest", string(data)}, " ")
}
