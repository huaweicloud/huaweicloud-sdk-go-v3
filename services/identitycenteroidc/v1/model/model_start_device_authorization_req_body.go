package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StartDeviceAuthorizationReqBody struct {

	// 在IAM身份中心注册的客户端的唯一标识符
	ClientId string `json:"client_id"`

	// 为客户端生成的秘密字符串。客户端将使用此字符串在后续调用中获得服务的身份验证
	ClientSecret string `json:"client_secret"`

	// 访问门户的URL
	StartUrl string `json:"start_url"`
}

func (o StartDeviceAuthorizationReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartDeviceAuthorizationReqBody struct{}"
	}

	return strings.Join([]string{"StartDeviceAuthorizationReqBody", string(data)}, " ")
}
