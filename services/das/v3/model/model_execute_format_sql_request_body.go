package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteFormatSqlRequestBody 格式化SQL请求体
type ExecuteFormatSqlRequestBody struct {

	// SQL语句
	SqlScript string `json:"sql_script"`
}

func (o ExecuteFormatSqlRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteFormatSqlRequestBody struct{}"
	}

	return strings.Join([]string{"ExecuteFormatSqlRequestBody", string(data)}, " ")
}
