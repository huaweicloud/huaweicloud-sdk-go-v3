package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetUserIdResponse Response Object
type GetUserIdResponse struct {

	// 身份源的全局唯一标识符（ID）
	IdentityStoreId *string `json:"identity_store_id,omitempty"`

	// 身份源中IAM身份中心用户的全局唯一标识符（ID）
	UserId         *string `json:"user_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o GetUserIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetUserIdResponse struct{}"
	}

	return strings.Join([]string{"GetUserIdResponse", string(data)}, " ")
}
