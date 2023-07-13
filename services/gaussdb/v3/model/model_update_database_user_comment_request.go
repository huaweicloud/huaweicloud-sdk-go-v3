package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDatabaseUserCommentRequest 修改数据库用户备注请求体。
type UpdateDatabaseUserCommentRequest struct {

	// 准备修改备注的数据库用户列表，列表最大长度为50。
	Users []UpdateDatabaseUserComment `json:"users"`
}

func (o UpdateDatabaseUserCommentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDatabaseUserCommentRequest struct{}"
	}

	return strings.Join([]string{"UpdateDatabaseUserCommentRequest", string(data)}, " ")
}
