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

	// 分段策略，多个策略之间用逗号分割。 > title:标题分割；separator:分隔符分割
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
