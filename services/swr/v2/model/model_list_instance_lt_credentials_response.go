package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceLtCredentialsResponse Response Object
type ListInstanceLtCredentialsResponse struct {

	// 长期访问凭据总数
	Total *int32 `json:"total,omitempty"`

	// 长期访问凭证列表
	AuthTokens     *[]AuthToken `json:"auth_tokens,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListInstanceLtCredentialsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceLtCredentialsResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceLtCredentialsResponse", string(data)}, " ")
}
