package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建数据库请求体。
type CreateGaussMySqlDatabaseRequestBody struct {
	Databases []CreateGaussMySqlDatabase `json:"databases"`
}

func (o CreateGaussMySqlDatabaseRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGaussMySqlDatabaseRequestBody struct{}"
	}

	return strings.Join([]string{"CreateGaussMySqlDatabaseRequestBody", string(data)}, " ")
}
