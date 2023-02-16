package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Tool struct {

	// 工具Id
	Id string `json:"id"`

	// 工具名称
	Name string `json:"name"`

	// 工具链接
	Url *string `json:"url,omitempty"`

	// 图标内容，Base64格式
	Icon *string `json:"icon,omitempty"`

	// 问题分类Id
	ProblemTypeId *string `json:"problem_type_id,omitempty"`

	// 业务类型Id
	BusinessTypeId *string `json:"business_type_id,omitempty"`

	// 产品类型Id
	ProductTypeId *string `json:"product_type_id,omitempty"`
}

func (o Tool) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tool struct{}"
	}

	return strings.Join([]string{"Tool", string(data)}, " ")
}
