package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExternalIdDto The identifier issued to this resource by an external identity provider.
type ExternalIdDto struct {

	// 外部身份提供商颁发给此资源的标识符
	Id string `json:"id"`

	// 外部标识符的颁发者
	Issuer string `json:"issuer"`
}

func (o ExternalIdDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExternalIdDto struct{}"
	}

	return strings.Join([]string{"ExternalIdDto", string(data)}, " ")
}
