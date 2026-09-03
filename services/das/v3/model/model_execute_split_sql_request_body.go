package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteSplitSqlRequestBody 拆分SQL请求体
type ExecuteSplitSqlRequestBody struct {

	// SQL语句
	SqlScript string `json:"sql_script"`
}

func (o ExecuteSplitSqlRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteSplitSqlRequestBody struct{}"
	}

	return strings.Join([]string{"ExecuteSplitSqlRequestBody", string(data)}, " ")
}
