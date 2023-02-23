package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PolicyItemDataMaskInfo struct {

	// 条件表达式
	ConditionExpr *string `json:"condition_expr,omitempty"`

	// 列掩码类型
	DataMaskType *string `json:"data_mask_type,omitempty"`

	// 列掩码表达式
	ValueExpr *string `json:"value_expr,omitempty"`
}

func (o PolicyItemDataMaskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyItemDataMaskInfo struct{}"
	}

	return strings.Join([]string{"PolicyItemDataMaskInfo", string(data)}, " ")
}
