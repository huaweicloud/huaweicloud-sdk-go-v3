package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowServerVncResponse Response Object
type ShowServerVncResponse struct {

	// 登录类型: * `novnc` * `vnc`
	Type *string `json:"type,omitempty"`

	// 远程登录控制台地址。
	Url *string `json:"url,omitempty"`

	// 登录协议。
	Protocol       *string `json:"protocol,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowServerVncResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerVncResponse struct{}"
	}

	return strings.Join([]string{"ShowServerVncResponse", string(data)}, " ")
}
