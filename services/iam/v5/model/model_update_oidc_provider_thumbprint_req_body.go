package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateOidcProviderThumbprintReqBody struct {

	// **参数解释**： OIDC 身份提供商的服务器证书指纹列表。  **约束限制**： 列表元素数量取值范围为[1,5]个，每个元素字符串长度为64。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
	Thumbprints []string `json:"thumbprints"`
}

func (o UpdateOidcProviderThumbprintReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOidcProviderThumbprintReqBody struct{}"
	}

	return strings.Join([]string{"UpdateOidcProviderThumbprintReqBody", string(data)}, " ")
}
