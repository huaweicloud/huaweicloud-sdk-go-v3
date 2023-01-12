package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FilterV2 struct {

	// 运算符。0：仅包含1：仅排除 此参数不携带或携带值为null时，不作为筛选条件。
	Operator int32 `json:"operator"`

	FilterFactor *FilterFactor `json:"filter_factor"`
}

func (o FilterV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FilterV2 struct{}"
	}

	return strings.Join([]string{"FilterV2", string(data)}, " ")
}
