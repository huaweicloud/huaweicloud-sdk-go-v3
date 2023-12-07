package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAvailableUpdatesResponse Response Object
type ShowAvailableUpdatesResponse struct {

	// Landing Zone基础配置是否可用。
	BaselineUpdateAvailable *bool `json:"baseline_update_available,omitempty"`

	// 控制策略是否可用。
	ControlUpdateAvailable *bool `json:"control_update_available,omitempty"`

	// Landing Zone是否可更新。
	LandingZoneUpdateAvailable *bool `json:"landing_zone_update_available,omitempty"`

	// Landing Zone当前最新的版本号。
	ServiceLandingZoneVersion *string `json:"service_landing_zone_version,omitempty"`

	// 用户当前的Landing Zone版本。
	UserLandingZoneVersion *string `json:"user_landing_zone_version,omitempty"`
	HttpStatusCode         int     `json:"-"`
}

func (o ShowAvailableUpdatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAvailableUpdatesResponse struct{}"
	}

	return strings.Join([]string{"ShowAvailableUpdatesResponse", string(data)}, " ")
}
