package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LoginBuiltInAccountResponse Response Object
type LoginBuiltInAccountResponse struct {

	// 数据库用户ID。
	DbUserId       *string `json:"db_user_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o LoginBuiltInAccountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoginBuiltInAccountResponse struct{}"
	}

	return strings.Join([]string{"LoginBuiltInAccountResponse", string(data)}, " ")
}
