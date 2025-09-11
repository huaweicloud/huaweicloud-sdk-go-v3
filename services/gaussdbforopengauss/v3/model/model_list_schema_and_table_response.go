package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSchemaAndTableResponse Response Object
type ListSchemaAndTableResponse struct {

	// **参数解释**: 数据库表信息列表。
	DatabaseTables *[]DatabaseSchemaTableResult `json:"database_tables,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListSchemaAndTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSchemaAndTableResponse struct{}"
	}

	return strings.Join([]string{"ListSchemaAndTableResponse", string(data)}, " ")
}
