package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TopSql top sql信息
type TopSql struct {

	// 采样时间
	SampleTime *string `json:"sample_time,omitempty"`

	// 个数
	Count *int32 `json:"count,omitempty"`

	// 数据库名
	DatabaseName *string `json:"database_name,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`

	// SQL语句
	SqlStatement *string `json:"sql_statement,omitempty"`

	// Query ID
	QueryId *string `json:"query_id,omitempty"`
}

func (o TopSql) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopSql struct{}"
	}

	return strings.Join([]string{"TopSql", string(data)}, " ")
}
