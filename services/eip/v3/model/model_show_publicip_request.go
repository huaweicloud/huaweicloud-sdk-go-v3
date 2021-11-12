package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowPublicipRequest struct {
	// 弹性公网IP的ID

	PublicipId string `json:"publicip_id"`
	// 显示，形式为\"fields=id&fields=owner&...\"  支持字段：id/project_id/ip_version/type/public_ip_address/public_ipv6_address/network_type/status/description/created_at/updated_at/vnic/bandwidth/associate_instance_type/associate_instance_id/lock_status/billing_info/tags/enterprise_project_id/publicip_pool_name/allow_share_bandwidth_types/alias/publicip_pool_id/public_border_group

	Fields *[]string `json:"fields,omitempty"`
}

func (o ShowPublicipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPublicipRequest struct{}"
	}

	return strings.Join([]string{"ShowPublicipRequest", string(data)}, " ")
}
