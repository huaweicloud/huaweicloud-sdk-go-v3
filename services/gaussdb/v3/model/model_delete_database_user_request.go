package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 删除数据库用户请求体。
type DeleteDatabaseUserRequest struct {

	// 准备删除的数据库用户列表，列表最大长度为50。
	Users []ListDeleteDatabaseUserRequest `json:"users"`
}

func (o DeleteDatabaseUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDatabaseUserRequest struct{}"
	}

	return strings.Join([]string{"DeleteDatabaseUserRequest", string(data)}, " ")
}
