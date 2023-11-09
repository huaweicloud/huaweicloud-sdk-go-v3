package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDesktopNetworkResponse Response Object
type ShowDesktopNetworkResponse struct {

	// 桌面名称
	ComputerName *string `json:"computer_name,omitempty"`

	// 桌面ID
	ComputerId *string `json:"computer_id,omitempty"`

	// 桌面网络信息
	NetworkInfos   *[]NetworkInfo `json:"network_infos,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowDesktopNetworkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDesktopNetworkResponse struct{}"
	}

	return strings.Join([]string{"ShowDesktopNetworkResponse", string(data)}, " ")
}
