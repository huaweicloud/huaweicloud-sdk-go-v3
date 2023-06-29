package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Product 产品详情
type Product struct {

	// 产品id
	Id *string `json:"id,omitempty"`

	// 产品名称
	Name string `json:"name"`

	// 产品描述
	Description *string `json:"description,omitempty"`

	// 产品属性值
	Attributes map[string]ProductAttributes `json:"attributes,omitempty"`

	// 产品所属账号的项目ID
	ProjectId string `json:"project_id"`

	// 产品创建时间
	CreatedAt *int32 `json:"created_at,omitempty"`

	// 产品标签
	Tags *[]Attributes `json:"tags,omitempty"`
}

func (o Product) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Product struct{}"
	}

	return strings.Join([]string{"Product", string(data)}, " ")
}
