package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProductMetadata 产品属性值基本信息
type ProductMetadata struct {

	// 产品属性值类型
	Type *string `json:"type,omitempty"`
}

func (o ProductMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductMetadata struct{}"
	}

	return strings.Join([]string{"ProductMetadata", string(data)}, " ")
}
