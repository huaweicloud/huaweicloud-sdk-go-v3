package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLandingZoneConfigurationResponse Response Object
type ShowLandingZoneConfigurationResponse struct {
	CommonConfiguration *CommonConfiguration `json:"common_configuration,omitempty"`

	LoggingConfiguration *LoggingConfiguration `json:"logging_configuration,omitempty"`

	OrganizationStructure *[]OrganizationStructureBaseLine `json:"organization_structure,omitempty"`

	// 纳管的区域
	Regions *[]RegionConfigurationList `json:"regions,omitempty"`

	// 管理员账号创建Identity Center用户所用邮箱
	IdentityStoreEmail *string `json:"identity_store_email,omitempty"`
	HttpStatusCode     int     `json:"-"`
}

func (o ShowLandingZoneConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLandingZoneConfigurationResponse struct{}"
	}

	return strings.Join([]string{"ShowLandingZoneConfigurationResponse", string(data)}, " ")
}
