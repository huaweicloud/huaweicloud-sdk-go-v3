package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatabaseSchemasResponse Response Object
type ListDatabaseSchemasResponse struct {

	// 列表中每个元素表示一个数据库schema。
	DatabaseSchemas *[]GaussDBforOpenGaussDatabaseForListSchema `json:"database_schemas,omitempty"`

	// 数据库schema总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDatabaseSchemasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabaseSchemasResponse struct{}"
	}

	return strings.Join([]string{"ListDatabaseSchemasResponse", string(data)}, " ")
}
