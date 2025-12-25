package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DocumentSegmentParam 修改文档分段参数。
type DocumentSegmentParam struct {

	// 分段类型 * 1: 自动分段 * 2: 手动分段
	SplitType int32 `json:"split_type"`

	// 分段长度。
	ChunkSize *int32 `json:"chunk_size,omitempty"`

	// 分段策略，如果添加多个策略用逗号隔开。 取值示例： - title：用标题把一段话分割为多个段落。 - separator：用分隔符把一段话分割为多个段落。
	ChunkType *string `json:"chunk_type,omitempty"`

	// 分隔符
	Separators *[]string `json:"separators,omitempty"`
}

func (o DocumentSegmentParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DocumentSegmentParam struct{}"
	}

	return strings.Join([]string{"DocumentSegmentParam", string(data)}, " ")
}
