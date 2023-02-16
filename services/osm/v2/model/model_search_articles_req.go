package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SearchArticlesReq struct {

	// 返回匹配度最高的数据条数
	Top *int32 `json:"top,omitempty"`

	// 产品类型Id
	ProductTypeId string `json:"product_type_id"`

	// 业务类型Id
	BusinessTypeId *string `json:"business_type_id,omitempty"`

	// 问题类型Id
	ProblemTypeId *string `json:"problem_type_id,omitempty"`
}

func (o SearchArticlesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchArticlesReq struct{}"
	}

	return strings.Join([]string{"SearchArticlesReq", string(data)}, " ")
}
