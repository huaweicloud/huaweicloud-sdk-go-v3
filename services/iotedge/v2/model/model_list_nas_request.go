package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListNasRequest struct {
	// NA名称搜索关键字

	Name *string `json:"name,omitempty"`
	// 查询的起始位置，取值范围为非负整数，默认为0

	Offset *int32 `json:"offset,omitempty"`
	// 每页记录数，取值范围为非负整数，默认值为10

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListNasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNasRequest struct{}"
	}

	return strings.Join([]string{"ListNasRequest", string(data)}, " ")
}
