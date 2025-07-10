package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLandingZoneStatusResponse Response Object
type ShowLandingZoneStatusResponse struct {

	// 部署的Landing Zone版本。
	DeployedVersion *string `json:"deployed_version,omitempty"`

	// Landing Zone的设置状态，包括进行中，已完成。
	LandingZoneStatus *string `json:"landing_zone_status,omitempty"`

	// Landing Zone的完成进度。
	PercentageComplete *int32 `json:"percentage_complete,omitempty"`

	// Landing Zone设置的详细进度信息。
	PercentageDetails *[]PercentageDetail `json:"percentage_details,omitempty"`

	// Landing Zone当前需要执行的动作。
	LandingZoneActionType *string `json:"landing_zone_action_type,omitempty"`

	// Landing Zone错误消息。
	Message *[]LandingZoneErrorMessage `json:"message,omitempty"`

	// 纳管的区域信息。
	Regions        *[]RegionConfigurationList `json:"regions,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ShowLandingZoneStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLandingZoneStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowLandingZoneStatusResponse", string(data)}, " ")
}
