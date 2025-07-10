package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDesktopNetworksResponse Response Object
type ShowDesktopNetworksResponse struct {

	// 桌面网络信息列表。
	Networks       *[]DesktopNetworkResult `json:"networks,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ShowDesktopNetworksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDesktopNetworksResponse struct{}"
	}

	return strings.Join([]string{"ShowDesktopNetworksResponse", string(data)}, " ")
}
