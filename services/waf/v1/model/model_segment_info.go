package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SegmentInfo 命中片段
type SegmentInfo struct {

	// 命中规则的具体片段
	Segment *string `json:"segment,omitempty"`

	// 命中的规则类别
	Category *string `json:"category,omitempty"`

	// 敏感词所在位置
	Position *[]int32 `json:"position,omitempty"`

	// 命中自定义词库的ID
	LexiconId *string `json:"lexicon_id,omitempty"`

	// 自定义词库是豁免还是违规
	LexiconType *int32 `json:"lexicon_type,omitempty"`
}

func (o SegmentInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SegmentInfo struct{}"
	}

	return strings.Join([]string{"SegmentInfo", string(data)}, " ")
}
