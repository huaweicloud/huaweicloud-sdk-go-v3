package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteFormatSqlResponse Response Object
type ExecuteFormatSqlResponse struct {

	// 格式化后的SQL语句
	FormatSql      *string `json:"format_sql,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteFormatSqlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteFormatSqlResponse struct{}"
	}

	return strings.Join([]string{"ExecuteFormatSqlResponse", string(data)}, " ")
}
