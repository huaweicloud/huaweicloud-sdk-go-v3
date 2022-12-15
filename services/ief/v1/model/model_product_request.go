package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 产品详情请求体
type ProductRequest struct {

	// 产品名称，允许输入小写字母，数字，中划线，不能以中划线开头或结尾，最大长度为26位
	Name string `json:"name"`

	// 产品描述
	Description *string `json:"description,omitempty"`

	// 产品属性值
	Attributes map[string]ProductAttributes `json:"attributes,omitempty"`

	// 产品标签
	Tags *[]Attributes `json:"tags,omitempty"`
}

func (o ProductRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductRequest struct{}"
	}

	return strings.Join([]string{"ProductRequest", string(data)}, " ")
}
