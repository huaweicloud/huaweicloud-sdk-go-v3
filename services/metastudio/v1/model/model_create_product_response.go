package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateProductResponse Response Object
type CreateProductResponse struct {

	// 商品ID
	ProductId *string `json:"product_id,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateProductResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProductResponse struct{}"
	}

	return strings.Join([]string{"CreateProductResponse", string(data)}, " ")
}
