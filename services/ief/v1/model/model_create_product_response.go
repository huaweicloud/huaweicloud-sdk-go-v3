package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateProductResponse Response Object
type CreateProductResponse struct {
	Product        *ProductResponse `json:"product,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o CreateProductResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProductResponse struct{}"
	}

	return strings.Join([]string{"CreateProductResponse", string(data)}, " ")
}
