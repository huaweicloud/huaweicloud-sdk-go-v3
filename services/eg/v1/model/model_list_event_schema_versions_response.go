package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListEventSchemaVersionsResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 本页数量
	Size *int32 `json:"size,omitempty"`

	// 对象列表
	Items          *[]CustomizeSchemaVersionItemInfo `json:"items,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o ListEventSchemaVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventSchemaVersionsResponse struct{}"
	}

	return strings.Join([]string{"ListEventSchemaVersionsResponse", string(data)}, " ")
}
