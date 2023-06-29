package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProductCreateRequest 产品参数
type ProductCreateRequest struct {
	Product *ProductRequest `json:"product"`
}

func (o ProductCreateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductCreateRequest struct{}"
	}

	return strings.Join([]string{"ProductCreateRequest", string(data)}, " ")
}
