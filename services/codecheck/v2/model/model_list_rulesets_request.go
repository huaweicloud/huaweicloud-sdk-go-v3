package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListRulesetsRequest struct {
	// 规则集类别  '0只查询系统规则集，1只查询当前用户自定义规则集，2只查询其他用户自定义规则集，'0,1,2'或''查所有'

	Category *string `json:"category,omitempty"`
	// 分页索引，偏移量

	Offset *int32 `json:"offset,omitempty"`
	// 每页显示的数量

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListRulesetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRulesetsRequest struct{}"
	}

	return strings.Join([]string{"ListRulesetsRequest", string(data)}, " ")
}
