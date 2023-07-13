package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEventTargetResponse Response Object
type ListEventTargetResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 本页数量
	Size *int32 `json:"size,omitempty"`

	// 对象列表
	Items          *[]CatalogTargetInfo `json:"items,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListEventTargetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventTargetResponse struct{}"
	}

	return strings.Join([]string{"ListEventTargetResponse", string(data)}, " ")
}
