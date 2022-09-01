package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QuerySqlPlanBody struct {

	// 数据库用户ID
	DbUserId *string `json:"db_user_id,omitempty" xml:"db_user_id"`

	// 数据库名称
	Database *string `json:"database,omitempty" xml:"database"`

	// SQL语句
	Sql *string `json:"sql,omitempty" xml:"sql"`
}

func (o QuerySqlPlanBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuerySqlPlanBody struct{}"
	}

	return strings.Join([]string{"QuerySqlPlanBody", string(data)}, " ")
}
