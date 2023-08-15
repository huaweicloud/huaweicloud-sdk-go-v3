package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SearchNoticesReq struct {

	// 返回匹配度最高的数据条数
	Top *int32 `json:"top,omitempty"`

	// 产品类型Id
	ProductTypeId string `json:"product_type_id"`
}

func (o SearchNoticesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchNoticesReq struct{}"
	}

	return strings.Join([]string{"SearchNoticesReq", string(data)}, " ")
}
