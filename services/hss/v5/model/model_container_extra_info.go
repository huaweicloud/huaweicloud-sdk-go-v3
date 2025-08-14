package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ContainerExtraInfo 沙箱额外配置，包括VPN引流沙箱、linux沙箱、RDP沙箱、mysql沙箱，其他沙箱默认为空
type ContainerExtraInfo struct {
	Openvpn *ContainerExtraOpenvpnInfo `json:"openvpn,omitempty"`

	Linux *ContainerExtraLinuxInfo `json:"linux,omitempty"`

	Rdp *ContainerExtraRdpInfo `json:"rdp,omitempty"`

	Mysql *ContainerExtraMysqlInfo `json:"mysql,omitempty"`
}

func (o ContainerExtraInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContainerExtraInfo struct{}"
	}

	return strings.Join([]string{"ContainerExtraInfo", string(data)}, " ")
}
