package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHdaUpgradeFlagRequest Request Object
type ShowHdaUpgradeFlagRequest struct {
}

func (o ShowHdaUpgradeFlagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHdaUpgradeFlagRequest struct{}"
	}

	return strings.Join([]string{"ShowHdaUpgradeFlagRequest", string(data)}, " ")
}
