package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRemoteConsoleAddressRequest Request Object
type ShowRemoteConsoleAddressRequest struct {

	// 桌面ID。
	DesktopId string `json:"desktop_id"`
}

func (o ShowRemoteConsoleAddressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRemoteConsoleAddressRequest struct{}"
	}

	return strings.Join([]string{"ShowRemoteConsoleAddressRequest", string(data)}, " ")
}
