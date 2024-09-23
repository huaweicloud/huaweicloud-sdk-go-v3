package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProductsResponse Response Object
type ListProductsResponse struct {

	// **参数解释**： 商品信息总数
	Count *int32 `json:"count,omitempty"`

	// 商品信息列表
	Products *[]ProductDetailInfo `json:"products,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListProductsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProductsResponse struct{}"
	}

	return strings.Join([]string{"ListProductsResponse", string(data)}, " ")
}
