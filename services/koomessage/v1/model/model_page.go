package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Page 分页信息。
type Page struct {

	// 偏移量，表示从此偏移量开始查询，offset大于等于0。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 总条数。
	Total *int32 `json:"total,omitempty"`
}

func (o Page) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Page struct{}"
	}

	return strings.Join([]string{"Page", string(data)}, " ")
}
