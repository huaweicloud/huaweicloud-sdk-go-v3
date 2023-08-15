package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetDatabasePasswordRequest 修改数据库用户密码请求体。
type ResetDatabasePasswordRequest struct {

	// 准备修改密码的数据库用户列表，列表最大长度为50。
	Users []ResetDatabasePassword `json:"users"`
}

func (o ResetDatabasePasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetDatabasePasswordRequest struct{}"
	}

	return strings.Join([]string{"ResetDatabasePasswordRequest", string(data)}, " ")
}
