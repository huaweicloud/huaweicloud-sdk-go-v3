package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ItemSource 数据的元信息，不同数据包含的字段可能不同。
type ItemSource struct {

	// 数据描述信息。
	Desc *string `json:"desc,omitempty"`

	// 数据自定义字符标签。
	CustomTags map[string]string `json:"custom_tags,omitempty"`

	// 数据自定义数值标签。
	CustomNumTags map[string]float64 `json:"custom_num_tags,omitempty"`

	// 数据关键词列表。
	Keywords *[]string `json:"keywords,omitempty"`
}

func (o ItemSource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ItemSource struct{}"
	}

	return strings.Join([]string{"ItemSource", string(data)}, " ")
}
