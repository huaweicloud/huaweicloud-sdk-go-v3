package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDesktopNetworkRequest Request Object
type ShowDesktopNetworkRequest struct {

	// 桌面ID。
	DesktopId string `json:"desktop_id"`
}

func (o ShowDesktopNetworkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDesktopNetworkRequest struct{}"
	}

	return strings.Join([]string{"ShowDesktopNetworkRequest", string(data)}, " ")
}
