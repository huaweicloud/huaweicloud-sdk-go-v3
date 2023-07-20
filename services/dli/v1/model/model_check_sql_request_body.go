package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckSqlRequestBody 检查SQL语法
type CheckSqlRequestBody struct {

	// 待执行的SQL语句。
	Sql string `json:"sql"`

	// SQL语句执行所在的数据库。
	Currentdb *string `json:"currentdb,omitempty"`
}

func (o CheckSqlRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckSqlRequestBody struct{}"
	}

	return strings.Join([]string{"CheckSqlRequestBody", string(data)}, " ")
}
