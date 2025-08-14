package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ContainerExtraOpenvpnInfo 沙箱类型，VPN引流沙箱专用
type ContainerExtraOpenvpnInfo struct {

	// 映射ip
	OutsideIp *string `json:"outside_ip,omitempty"`

	// 映射端口
	OutsidePort *string `json:"outside_port,omitempty"`
}

func (o ContainerExtraOpenvpnInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContainerExtraOpenvpnInfo struct{}"
	}

	return strings.Join([]string{"ContainerExtraOpenvpnInfo", string(data)}, " ")
}
