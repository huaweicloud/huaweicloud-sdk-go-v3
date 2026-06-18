package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveClientIdFromOidcProviderV5Request Request Object
type RemoveClientIdFromOidcProviderV5Request struct {

	// 提供商ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	ProviderId string `json:"provider_id"`

	Body *RemoveClientIdFromOidcProviderReqBody `json:"body,omitempty"`
}

func (o RemoveClientIdFromOidcProviderV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveClientIdFromOidcProviderV5Request struct{}"
	}

	return strings.Join([]string{"RemoveClientIdFromOidcProviderV5Request", string(data)}, " ")
}
