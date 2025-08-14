package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IdpSamlConfig struct {

	// 身份提供商发布者标识
	EntityId string `json:"entity_id"`

	// 身份提供商登录链接
	LoginUrl string `json:"login_url"`

	// 是否要求SAML请求签名验证
	WantRequestSigned bool `json:"want_request_signed"`
}

func (o IdpSamlConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdpSamlConfig struct{}"
	}

	return strings.Join([]string{"IdpSamlConfig", string(data)}, " ")
}
