package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSharerProductsResponse Response Object
type ListSharerProductsResponse struct {

	// 产品列表。
	Products *[]SharerProductInfo `json:"products,omitempty"`

	// 对象总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSharerProductsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSharerProductsResponse struct{}"
	}

	return strings.Join([]string{"ListSharerProductsResponse", string(data)}, " ")
}
