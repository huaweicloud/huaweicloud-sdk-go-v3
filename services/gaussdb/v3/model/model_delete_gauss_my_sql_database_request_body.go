package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGaussMySqlDatabaseRequestBody 删除数据库请求体。
type DeleteGaussMySqlDatabaseRequestBody struct {

	// 准备删除的数据库列表，列表最大长度为50。
	Databases []string `json:"databases"`
}

func (o DeleteGaussMySqlDatabaseRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGaussMySqlDatabaseRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteGaussMySqlDatabaseRequestBody", string(data)}, " ")
}
