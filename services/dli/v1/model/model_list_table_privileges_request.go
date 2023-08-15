package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTablePrivilegesRequest Request Object
type ListTablePrivilegesRequest struct {

	// 被查询的数据库名称。
	DatabaseName string `json:"database_name"`

	// 被查询的表名称。
	TableName string `json:"table_name"`

	// 被查询的用户名称。
	UserName string `json:"user_name"`
}

func (o ListTablePrivilegesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTablePrivilegesRequest struct{}"
	}

	return strings.Join([]string{"ListTablePrivilegesRequest", string(data)}, " ")
}
