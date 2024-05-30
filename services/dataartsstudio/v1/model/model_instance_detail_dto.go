package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceDetailDto 集群详情信息
type InstanceDetailDto struct {

	// 集群ID。
	Id *string `json:"id,omitempty"`

	// 集群名称。
	Name *string `json:"name,omitempty"`

	// 集群描述信息。
	Description *string `json:"description,omitempty"`

	// 公网IP地址。
	ExternalAddress *string `json:"external_address,omitempty"`

	// 内网IPv4地址。
	IntranetAddress *string `json:"intranet_address,omitempty"`

	// 内网IPv6地址。
	IntranetAddressIpv6 *string `json:"intranet_address_ipv6,omitempty"`

	// 公网域名ID。
	PublicZoneId *string `json:"public_zone_id,omitempty"`

	// 公网域名名称。
	PublicZoneName *string `json:"public_zone_name,omitempty"`

	// 内网域名ID。
	PrivateZoneId *string `json:"private_zone_id,omitempty"`

	// 内网域名名称。
	PrivateZoneName *string `json:"private_zone_name,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 创建时间。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 创建人。
	CreateUser *string `json:"create_user,omitempty"`

	// 当前工作空间已发布的API数量。
	CurrentNamespacePublishApiNum *int32 `json:"current_namespace_publish_api_num,omitempty"`

	// 所有工作空间已发布的API数量。
	AllNamespacePublishApiNum *int32 `json:"all_namespace_publish_api_num,omitempty"`

	// 集群API总配额。
	ApiPublishableNum *int32 `json:"api_publishable_num,omitempty"`

	// 集群是否可以删除。
	Deletable *bool `json:"deletable,omitempty"`

	// 集群计费状态，NO_CHARGE：未计费、CHARGED：已计费，GRACE：宽限期、RETENTION：保留期。
	ChargeStatus *string `json:"charge_status,omitempty"`

	// 订单ID。
	OrderId *string `json:"order_id,omitempty"`

	// 订单类型，PERIOD：包周期、ON_DEMAND：按需。
	OrderType *string `json:"order_type,omitempty"`

	// 集群订购周期类型。
	PeriodType *string `json:"period_type,omitempty"`

	// 集群状态。
	InstanceStatus *string `json:"instance_status,omitempty"`

	// 节点数量。
	NodeNum *int32 `json:"node_num,omitempty"`

	Flavor *FlavorDto `json:"flavor,omitempty"`

	// 集群版本号。
	GatewayVersion *string `json:"gateway_version,omitempty"`

	// 集群所在可用区编码。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 集群所在可用区名称。
	AvailabilityZoneName *string `json:"availability_zone_name,omitempty"`

	// 集群所在虚拟私有云ID。
	VpcId *string `json:"vpc_id,omitempty"`

	// 集群所在子网ID。
	SubnetId *string `json:"subnet_id,omitempty"`

	// 集群所在安全组ID。
	SecurityGroupId *string `json:"security_group_id,omitempty"`

	// 集群节点列表。
	Nodes *[]InstanceNodeDto `json:"nodes,omitempty"`
}

func (o InstanceDetailDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceDetailDto struct{}"
	}

	return strings.Join([]string{"InstanceDetailDto", string(data)}, " ")
}
