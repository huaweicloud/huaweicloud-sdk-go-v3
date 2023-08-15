package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Article struct {

	// 案例Id
	Id *string `json:"id,omitempty"`

	// 案例编码
	Code *string `json:"code,omitempty"`

	// 案例名称
	Name *string `json:"name,omitempty"`

	// 案例链接
	Url *string `json:"url,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 产品类型Id
	ProductTypeId *string `json:"product_type_id,omitempty"`

	// 业务类型Id
	BusinessTypeId *string `json:"business_type_id,omitempty"`

	// 问题类型Id
	ProblemTypeId *string `json:"problem_type_id,omitempty"`
}

func (o Article) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Article struct{}"
	}

	return strings.Join([]string{"Article", string(data)}, " ")
}
