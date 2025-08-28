package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DocumentSegmentInfo 文档分段信息
type DocumentSegmentInfo struct {

	// 分段序号
	TextIndex *int32 `json:"text_index,omitempty"`

	// 分段文本ID
	Id *string `json:"id,omitempty"`

	// 分段文本标题
	Title *string `json:"title,omitempty"`

	// 分段文本内容
	Text *string `json:"text,omitempty"`
}

func (o DocumentSegmentInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DocumentSegmentInfo struct{}"
	}

	return strings.Join([]string{"DocumentSegmentInfo", string(data)}, " ")
}
