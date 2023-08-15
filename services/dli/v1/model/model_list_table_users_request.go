package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTableUsersRequest Request Object
type ListTableUsersRequest struct {

	// 被查询的数据库名称。
	DatabaseName string `json:"database_name"`

	// 被查询的表名称。
	TableName string `json:"table_name"`
}

func (o ListTableUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTableUsersRequest struct{}"
	}

	return strings.Join([]string{"ListTableUsersRequest", string(data)}, " ")
}
