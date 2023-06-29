package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EncodeServerAccessInfo 编码服务访问信息。
type EncodeServerAccessInfo struct {

	// 编码服务监听端口
	ListenPort *int32 `json:"listen_port,omitempty"`

	// 编码服务访问的公网IP（过期）
	AccessIp *string `json:"access_ip,omitempty"`

	// 编码服务访问的内网IP（过期）
	IntranetIp *string `json:"intranet_ip,omitempty"`

	// 编码服务访问的公网IP（新增）
	PublicIp *string `json:"public_ip,omitempty"`

	// 编码服务访问的内网IP（新增）
	ServerIp *string `json:"server_ip,omitempty"`

	// 编码服务公网的访问端口
	AccessPort *int32 `json:"access_port,omitempty"`

	// 编码服务的端口类型，取值如下： - adb：云手机的ADB端口 - vnc：云手机的VNC端口 - cph_app_server：云游戏客户端接入端 - cph_h5_server：云游戏H5 web网页接入端口 - 其他值：用户自定义端口
	Type *string `json:"type,omitempty"`
}

func (o EncodeServerAccessInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EncodeServerAccessInfo struct{}"
	}

	return strings.Join([]string{"EncodeServerAccessInfo", string(data)}, " ")
}
