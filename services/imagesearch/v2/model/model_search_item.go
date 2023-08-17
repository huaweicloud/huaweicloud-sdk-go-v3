package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchItem 搜索结果参数结构体。
type SearchItem struct {

	// 数据唯一ID。
	Id *string `json:"id,omitempty"`

	// 数据匹配分数。
	Score *float64 `json:"score,omitempty"`

	Source *ItemSource `json:"source,omitempty"`
}

func (o SearchItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchItem struct{}"
	}

	return strings.Join([]string{"SearchItem", string(data)}, " ")
}
