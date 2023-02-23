package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PolicyItemRowFilterInfo struct {

	// 行过滤表达式
	FilterExpr *string `json:"filter_expr,omitempty"`
}

func (o PolicyItemRowFilterInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyItemRowFilterInfo struct{}"
	}

	return strings.Join([]string{"PolicyItemRowFilterInfo", string(data)}, " ")
}
