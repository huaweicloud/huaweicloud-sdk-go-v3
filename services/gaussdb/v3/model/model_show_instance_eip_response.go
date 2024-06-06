package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceEipResponse Response Object
type ShowInstanceEipResponse struct {

	// 是否能访问公网
	CanEnablePublicAccess *bool `json:"can_enable_public_access,omitempty"`

	// 弹性公网ID。
	Id *string `json:"id,omitempty"`

	// 弹性公网IP的网络类型。
	Type *string `json:"type,omitempty"`

	// 端口ID。
	PortId *string `json:"port_id,omitempty"`

	// 弹性公网IP地址。
	PublicIpAddress *string `json:"public_ip_address,omitempty"`

	// 私网IP地址。
	PrivateIpAddress *string `json:"private_ip_address,omitempty"`

	// 弹性公网IP地址。
	Status *string `json:"status,omitempty"`

	// 租户ID。
	TenantId *string `json:"tenant_id,omitempty"`

	// 创建时间。
	CreateTime *string `json:"create_time,omitempty"`

	// 带宽ID。
	BandwidthId *string `json:"bandwidth_id,omitempty"`

	// 带宽名称。
	BandwidthName *string `json:"bandwidth_name,omitempty"`

	// 带宽大小。
	BandwidthSize *string `json:"bandwidth_size,omitempty"`

	// 带宽类型。枚举值：PER 和WHOLE。
	BandwidthShareType *string `json:"bandwidth_share_type,omitempty"`

	// 额外参数，包括订单id、产品id等信息。如果profile不为空，说明是包周期的弹性公网IP。
	Profile        *interface{} `json:"profile,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowInstanceEipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceEipResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceEipResponse", string(data)}, " ")
}
