package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListPostgresqlDatabasesRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty" xml:"X-Language"`

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 分页页码，从1开始。
	Page int32 `json:"page" xml:"page"`

	// 每页数据条数。取值范围[1, 100]。
	Limit int32 `json:"limit" xml:"limit"`
}

func (o ListPostgresqlDatabasesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPostgresqlDatabasesRequest struct{}"
	}

	return strings.Join([]string{"ListPostgresqlDatabasesRequest", string(data)}, " ")
}
