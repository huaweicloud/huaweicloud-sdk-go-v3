package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProductAttributes 产品属性
type ProductAttributes struct {

	// 产品属性值
	Value *string `json:"value,omitempty"`

	Metadata *ProductMetadata `json:"metadata,omitempty"`
}

func (o ProductAttributes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductAttributes struct{}"
	}

	return strings.Join([]string{"ProductAttributes", string(data)}, " ")
}
