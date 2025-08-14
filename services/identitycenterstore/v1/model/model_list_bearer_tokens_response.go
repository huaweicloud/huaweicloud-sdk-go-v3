package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBearerTokensResponse Response Object
type ListBearerTokensResponse struct {

	// 访问令牌列表
	BearerTokens   *[]BearerToken `json:"bearer_tokens,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListBearerTokensResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBearerTokensResponse struct{}"
	}

	return strings.Join([]string{"ListBearerTokensResponse", string(data)}, " ")
}
