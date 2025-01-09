package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDesktopNetworksRequest Request Object
type ShowDesktopNetworksRequest struct {

	// 桌面id列表，最小为1，最大为100。
	DesktopIds *[]string `json:"desktop_ids,omitempty"`
}

func (o ShowDesktopNetworksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDesktopNetworksRequest struct{}"
	}

	return strings.Join([]string{"ShowDesktopNetworksRequest", string(data)}, " ")
}
