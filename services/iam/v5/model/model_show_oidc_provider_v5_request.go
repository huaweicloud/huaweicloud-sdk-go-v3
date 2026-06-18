package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOidcProviderV5Request Request Object
type ShowOidcProviderV5Request struct {

	// 提供商ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	ProviderId string `json:"provider_id"`
}

func (o ShowOidcProviderV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOidcProviderV5Request struct{}"
	}

	return strings.Join([]string{"ShowOidcProviderV5Request", string(data)}, " ")
}
