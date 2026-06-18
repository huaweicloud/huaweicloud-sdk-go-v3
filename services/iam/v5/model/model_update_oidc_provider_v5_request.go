package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateOidcProviderV5Request Request Object
type UpdateOidcProviderV5Request struct {

	// 提供商ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	ProviderId string `json:"provider_id"`

	Body *UpdateOidcProviderReqBody `json:"body,omitempty"`
}

func (o UpdateOidcProviderV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOidcProviderV5Request struct{}"
	}

	return strings.Join([]string{"UpdateOidcProviderV5Request", string(data)}, " ")
}
