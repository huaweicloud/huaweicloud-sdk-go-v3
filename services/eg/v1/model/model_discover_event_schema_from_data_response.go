package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DiscoverEventSchemaFromDataResponse Response Object
type DiscoverEventSchemaFromDataResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 本页数量
	Size *int32 `json:"size,omitempty"`

	// 对象列表
	Items          *[]CustomizeSchemaItemInfo `json:"items,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o DiscoverEventSchemaFromDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiscoverEventSchemaFromDataResponse struct{}"
	}

	return strings.Join([]string{"DiscoverEventSchemaFromDataResponse", string(data)}, " ")
}
