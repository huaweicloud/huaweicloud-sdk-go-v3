package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLandingZoneConfigurationResponse Response Object
type ShowLandingZoneConfigurationResponse struct {
	CommonConfiguration *CommonConfiguration `json:"common_configuration,omitempty"`

	LoggingConfiguration *LoggingConfiguration `json:"logging_configuration,omitempty"`

	OrganizationStructure *[]OrganizationStructureBaseLineRsp `json:"organization_structure,omitempty"`

	// 纳管的区域信息。
	Regions        *[]RegionConfigurationList `json:"regions,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ShowLandingZoneConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLandingZoneConfigurationResponse struct{}"
	}

	return strings.Join([]string{"ShowLandingZoneConfigurationResponse", string(data)}, " ")
}
