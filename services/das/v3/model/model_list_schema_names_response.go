package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSchemaNamesResponse Response Object
type ListSchemaNamesResponse struct {

	// 数据库对象信息列表
	Data *[]SchemaList `json:"data,omitempty"`

	// 列表大小
	Total          *int64 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSchemaNamesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSchemaNamesResponse struct{}"
	}

	return strings.Join([]string{"ListSchemaNamesResponse", string(data)}, " ")
}
