package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListEnvironmentsRequest struct {
	// 起始偏移量，表示从此偏移量开始查询， offset大于等于0

	Offset int64 `json:"offset"`
	// 每页显示的条目数量,最大支持200条

	Limit int64 `json:"limit"`
}

func (o ListEnvironmentsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnvironmentsRequest struct{}"
	}

	return strings.Join([]string{"ListEnvironmentsRequest", string(data)}, " ")
}
