package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DesktopNetworkResult 桌面网络查询结果
type DesktopNetworkResult struct {

	// 桌面名称
	ComputerName *string `json:"computer_name,omitempty"`

	// 桌面ID
	ComputerId *string `json:"computer_id,omitempty"`

	// 桌面IP
	ComputerIp *string `json:"computer_ip,omitempty"`

	// 桌面网络信息
	NetworkInfos *[]NetworkInfo `json:"network_infos,omitempty"`
}

func (o DesktopNetworkResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DesktopNetworkResult struct{}"
	}

	return strings.Join([]string{"DesktopNetworkResult", string(data)}, " ")
}
