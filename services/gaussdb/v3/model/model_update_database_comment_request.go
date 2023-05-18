package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 修改数据库备注请求体。
type UpdateDatabaseCommentRequest struct {

	// 准备修改备注的数据库列表，列表最大长度为50。
	Databases []UpdateDatabaseComment `json:"databases"`
}

func (o UpdateDatabaseCommentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDatabaseCommentRequest struct{}"
	}

	return strings.Join([]string{"UpdateDatabaseCommentRequest", string(data)}, " ")
}
