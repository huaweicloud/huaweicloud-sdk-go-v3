package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateOidcProviderThumbprintV5Request Request Object
type UpdateOidcProviderThumbprintV5Request struct {

	// 提供商ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	ProviderId string `json:"provider_id"`

	Body *UpdateOidcProviderThumbprintReqBody `json:"body,omitempty"`
}

func (o UpdateOidcProviderThumbprintV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOidcProviderThumbprintV5Request struct{}"
	}

	return strings.Join([]string{"UpdateOidcProviderThumbprintV5Request", string(data)}, " ")
}
