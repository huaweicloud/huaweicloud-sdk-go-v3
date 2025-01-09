package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRemoteConsoleAddressResponse Response Object
type ShowRemoteConsoleAddressResponse struct {

	// 登录类型
	Type *string `json:"type,omitempty"`

	// 远程登录控制台地址
	Url *string `json:"url,omitempty"`

	// 登录协议
	Protocol       *string `json:"protocol,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowRemoteConsoleAddressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRemoteConsoleAddressResponse struct{}"
	}

	return strings.Join([]string{"ShowRemoteConsoleAddressResponse", string(data)}, " ")
}
