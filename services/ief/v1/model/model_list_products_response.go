package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProductsResponse Response Object
type ListProductsResponse struct {

	// 产品作业数目
	Count *int32 `json:"count,omitempty"`

	// 产品作业详情
	Products       *[]Product `json:"products,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListProductsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProductsResponse struct{}"
	}

	return strings.Join([]string{"ListProductsResponse", string(data)}, " ")
}
