package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ContainerExtraRdpInfo 沙箱类型，RDP沙箱专用
type ContainerExtraRdpInfo struct {

	// 协议类型: - 0 : 协议模拟 - 1 : 系统模拟
	ProtoEnv *string `json:"proto_env,omitempty"`

	// 系统类型，系统模拟时使用: - win7 : win 7 - win8 : win 8 - win10 : win 10 - win-server2012 : win-server 2012 - win-server2016 : win-server 2016
	System *string `json:"system,omitempty"`
}

func (o ContainerExtraRdpInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContainerExtraRdpInfo struct{}"
	}

	return strings.Join([]string{"ContainerExtraRdpInfo", string(data)}, " ")
}
