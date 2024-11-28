package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLivePlatformProductsResponse Response Object
type ListLivePlatformProductsResponse struct {

	// 商品总数。性能考虑仅在offset=0时返回。
	Count *int32 `json:"count,omitempty"`

	// 任务ID
	Products *[]PlatformProductInfo `json:"products,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListLivePlatformProductsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLivePlatformProductsResponse struct{}"
	}

	return strings.Join([]string{"ListLivePlatformProductsResponse", string(data)}, " ")
}
