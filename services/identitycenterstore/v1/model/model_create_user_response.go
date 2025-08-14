package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateUserResponse Response Object
type CreateUserResponse struct {

	// 身份源的全局唯一标识符（ID）
	IdentityStoreId *string `json:"identity_store_id,omitempty"`

	// 身份源中IdentityCenter用户的全局唯一标识符（ID）
	UserId *string `json:"user_id,omitempty"`

	// 用于初始化密码的一次性密码
	Password       *string `json:"password,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUserResponse struct{}"
	}

	return strings.Join([]string{"CreateUserResponse", string(data)}, " ")
}
