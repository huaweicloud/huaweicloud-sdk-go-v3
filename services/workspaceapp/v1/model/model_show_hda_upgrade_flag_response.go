package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHdaUpgradeFlagResponse Response Object
type ShowHdaUpgradeFlagResponse struct {

	// hda是否需要升级标识。0为无需提示升级, 1为需提示升级
	UpgradeFlag    *int32 `json:"upgrade_flag,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowHdaUpgradeFlagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHdaUpgradeFlagResponse struct{}"
	}

	return strings.Join([]string{"ShowHdaUpgradeFlagResponse", string(data)}, " ")
}
