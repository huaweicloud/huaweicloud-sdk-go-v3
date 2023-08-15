package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProductsResponse Response Object
type ListProductsResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 本次返回数量
	Size *int32 `json:"size,omitempty"`

	// 产品列表
	Items          *[]Product `json:"items,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListProductsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProductsResponse struct{}"
	}

	return strings.Join([]string{"ListProductsResponse", string(data)}, " ")
}
