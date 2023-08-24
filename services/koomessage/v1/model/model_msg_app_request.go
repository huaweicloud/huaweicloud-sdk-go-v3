package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MsgAppRequest 创建短信应用请求体。
type MsgAppRequest struct {

	// 应用名称。
	AppName string `json:"app_name"`

	// 上行回调地址。支持通信协议HTTPS/HTTP。
	UpLinkAddr *string `json:"up_link_addr,omitempty"`
}

func (o MsgAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MsgAppRequest struct{}"
	}

	return strings.Join([]string{"MsgAppRequest", string(data)}, " ")
}
