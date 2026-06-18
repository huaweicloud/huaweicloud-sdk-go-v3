package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddClientIdToOidcProviderV5Request Request Object
type AddClientIdToOidcProviderV5Request struct {

	// 提供商ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	ProviderId string `json:"provider_id"`

	Body *AddClientIdToOidcProviderReqBody `json:"body,omitempty"`
}

func (o AddClientIdToOidcProviderV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddClientIdToOidcProviderV5Request struct{}"
	}

	return strings.Join([]string{"AddClientIdToOidcProviderV5Request", string(data)}, " ")
}
