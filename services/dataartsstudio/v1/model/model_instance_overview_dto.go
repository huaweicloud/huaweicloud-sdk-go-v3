package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceOverviewDto 集群概览信息
type InstanceOverviewDto struct {

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
}

func (o InstanceOverviewDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceOverviewDto struct{}"
	}

	return strings.Join([]string{"InstanceOverviewDto", string(data)}, " ")
}
