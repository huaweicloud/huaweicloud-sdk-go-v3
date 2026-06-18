package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RemoveClientIdFromOidcProviderReqBody struct {

	// **参数解释**： 要从 OIDC 提供商移除的客户端 ID。  **约束限制**： 长度范围为[1,255]。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
	ClientId string `json:"client_id"`
}

func (o RemoveClientIdFromOidcProviderReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveClientIdFromOidcProviderReqBody struct{}"
	}

	return strings.Join([]string{"RemoveClientIdFromOidcProviderReqBody", string(data)}, " ")
}
