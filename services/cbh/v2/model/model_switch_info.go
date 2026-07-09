package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchInfo 服务功能支持开关信息。
type SwitchInfo struct {

	// 是否支持unibuy。 - true：是 - false：否
	IsSupportUnibuy *bool `json:"is_support_unibuy,omitempty"`

	// 是否支持浮动IPv6。 - true：是 - false：否
	IsSupportFloatIpv6 *bool `json:"is_support_float_ipv6,omitempty"`

	// 是否支持管理员登录。 - true：是 - false：否
	IsSupportAdminLogin *bool `json:"is_support_admin_login,omitempty"`

	// 是否支持更新HA。 - true：是 - false：否
	IsSupportUpdateHa *bool `json:"is_support_update_ha,omitempty"`

	// 是否支持TMS。 - true：是 - false：否
	IsSupportTms *bool `json:"is_support_tms,omitempty"`

	// 是否支持EPS。 - true：是 - false：否
	IsSupportEps *bool `json:"is_support_eps,omitempty"`

	// 是否支持IAM登录。 - true：是 - false：否
	IsSupportIamLogin *bool `json:"is_support_iam_login,omitempty"`

	// 是否支持IPv6。 - true：是 - false：否
	IsSupportIpv6 *bool `json:"is_support_ipv6,omitempty"`

	// 是否支持HA。 - true：是 - false：否
	IsSupportHa *bool `json:"is_support_ha,omitempty"`

	// 是否支持重置。 - true：是 - false：否
	IsSupportReset *bool `json:"is_support_reset,omitempty"`

	// 是否支持升级实例。 - true：是 - false：否
	IsSupportUpgradeInstance *bool `json:"is_support_upgrade_instance,omitempty"`

	// 是否支持更改安全组。 - true：是 - false：否
	IsSupportChangeSecurityGroup *bool `json:"is_support_change_security_group,omitempty"`

	// 是否支持手动IP。 - true：是 - false：否
	IsSupportManuallyIp *bool `json:"is_support_manually_ip,omitempty"`

	// 是否支持容量扩展。 - true：是 - false：否
	IsSupportCapacityExpantion *bool `json:"is_support_capacity_expantion,omitempty"`

	// 是否支持HA扩展。 - true：是 - false：否
	IsSupportHaExpantion *bool `json:"is_support_ha_expantion,omitempty"`

	// 是否支持代理授权。 - true：是 - false：否
	IsSupportAgencyAuthorize *bool `json:"is_support_agency_authorize,omitempty"`

	// 是否支持更改VPC。 - true：是 - false：否
	IsSupportChangeVpc *bool `json:"is_support_change_vpc,omitempty"`

	// 是否支持集群。 - true：是 - false：否
	IsSupportCluster *bool `json:"is_support_cluster,omitempty"`

	// 是否支持按需。 - true：是 - false：否
	IsSupportOndemand *bool `json:"is_support_ondemand,omitempty"`

	// 是否支持周期。 - true：是 - false：否
	IsSupportPeriod *bool `json:"is_support_period,omitempty"`
}

func (o SwitchInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchInfo struct{}"
	}

	return strings.Join([]string{"SwitchInfo", string(data)}, " ")
}
