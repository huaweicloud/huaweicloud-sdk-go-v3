package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAvailableUpdatesResponse Response Object
type ShowAvailableUpdatesResponse struct {

	// 用户当前的Landing Zone版本是否为最新版本。
	BaselineUpdateAvailable *bool `json:"baseline_update_available,omitempty"`

	// 当前账号下是否有新的控制策略。
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
