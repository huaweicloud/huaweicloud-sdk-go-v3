package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatabaseUsersRequest Request Object
type ListDatabaseUsersRequest struct {

	// 被查询的数据库名称。
	DatabaseName string `json:"database_name"`
}

func (o ListDatabaseUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabaseUsersRequest struct{}"
	}

	return strings.Join([]string{"ListDatabaseUsersRequest", string(data)}, " ")
}
