package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListSqlserverDatabasesRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty" xml:"X-Language"`

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 分页页码，从1开始。
	Page int32 `json:"page" xml:"page"`

	// 每页数据条数。取值范围[1, 100]。
	Limit int32 `json:"limit" xml:"limit"`

	// 数据库名。当指定该参数时，page和limit参数需要传入但不生效。
	DbName *string `json:"db-name,omitempty" xml:"db-name"`
}

func (o ListSqlserverDatabasesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlserverDatabasesRequest struct{}"
	}

	return strings.Join([]string{"ListSqlserverDatabasesRequest", string(data)}, " ")
}
