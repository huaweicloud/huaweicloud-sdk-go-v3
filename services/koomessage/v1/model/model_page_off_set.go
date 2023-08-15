package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PageOffSet struct {

	// 偏移量，表示从此偏移量开始查询，offset大于等于0。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 总量。
	Total *int32 `json:"total,omitempty"`
}

func (o PageOffSet) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageOffSet struct{}"
	}

	return strings.Join([]string{"PageOffSet", string(data)}, " ")
}
