package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSchemasResponse Response Object
type ListSchemasResponse struct {

	// 当前数据连接schema记录数
	TotalCount *int32 `json:"total_count,omitempty"`

	// 数据连接id
	DwId *string `json:"dw_id,omitempty"`

	// 数据库名称
	Database *string `json:"database,omitempty"`

	// schema列表
	Schemas        *[]SchemasList `json:"schemas,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListSchemasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSchemasResponse struct{}"
	}

	return strings.Join([]string{"ListSchemasResponse", string(data)}, " ")
}
