package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBearerTokensRequest Request Object
type ListBearerTokensRequest struct {

	// 身份源的全局唯一标识符（ID）
	IdentityStoreId string `json:"identity_store_id"`

	// 自动预置配置的全局唯一标识符（ID）
	TenantId string `json:"tenant_id"`

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`
}

func (o ListBearerTokensRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBearerTokensRequest struct{}"
	}

	return strings.Join([]string{"ListBearerTokensRequest", string(data)}, " ")
}
