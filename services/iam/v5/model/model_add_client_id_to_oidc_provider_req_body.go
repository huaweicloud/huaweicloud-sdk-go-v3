package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddClientIdToOidcProviderReqBody struct {

	// **参数解释**： 要添加到 OIDC 提供商的客户端 ID。  **约束限制**： 长度范围为[1,255]。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
	ClientId string `json:"client_id"`
}

func (o AddClientIdToOidcProviderReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddClientIdToOidcProviderReqBody struct{}"
	}

	return strings.Join([]string{"AddClientIdToOidcProviderReqBody", string(data)}, " ")
}
