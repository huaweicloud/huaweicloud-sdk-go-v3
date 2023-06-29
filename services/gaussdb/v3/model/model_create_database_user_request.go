package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDatabaseUserRequest 创建数据库用户的请求信息。
type CreateDatabaseUserRequest struct {
	Users []CreateDatabaseUserList `json:"users"`
}

func (o CreateDatabaseUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatabaseUserRequest struct{}"
	}

	return strings.Join([]string{"CreateDatabaseUserRequest", string(data)}, " ")
}
