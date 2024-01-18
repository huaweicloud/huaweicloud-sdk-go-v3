package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetupLandingZoneReqBody The request body of the CreateResourceShare operation.
type SetupLandingZoneReqBody struct {

	// 管理员纳管账号创建Identity Center用户所用邮箱。
	IdentityStoreEmail string `json:"identity_store_email"`

	// 主区域。
	HomeRegion string `json:"home_region"`

	// 设置Landing Zone的类型。包括CREATE、REPAIR以及UPDATE。
	SetupLandingZoneActionType string `json:"setup_landing_zone_action_type"`

	// 当前纳管账号纳管的区域。
	RegionConfigurationList []RegionConfigurationList `json:"region_configuration_list"`

	// 基础环境的纳管账号体系。
	OrganizationStructure []OrganizationStructureBaseLine `json:"organization_structure"`

	// 是否允许区域拒绝，默认false。
	DenyUngovernedRegions *bool `json:"deny_ungoverned_regions,omitempty"`

	// 是否允许设置组织汇聚。
	CloudTrailType *bool `json:"cloud_trail_type,omitempty"`

	// 加密字段。
	KmsKeyId *string `json:"kms_key_id,omitempty"`

	LoggingConfiguration *SetupLandingZoneReqBodyLoggingConfiguration `json:"logging_configuration"`
}

func (o SetupLandingZoneReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetupLandingZoneReqBody struct{}"
	}

	return strings.Join([]string{"SetupLandingZoneReqBody", string(data)}, " ")
}
