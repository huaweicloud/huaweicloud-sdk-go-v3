package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEventSchemaResponse Response Object
type ListEventSchemaResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 本页数量
	Size *int32 `json:"size,omitempty"`

	// 对象列表
	Items          *[]CustomizeSchemaItemInfo `json:"items,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListEventSchemaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventSchemaResponse struct{}"
	}

	return strings.Join([]string{"ListEventSchemaResponse", string(data)}, " ")
}
