package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SegmentResult struct {

	// 命中的风险片段。
	Segment *string `json:"segment,omitempty"`

	// 命中的自定义词库名称。  命中自定义词库时，才会返回当前字段。
	GlossaryName *string `json:"glossary_name,omitempty"`

	// 命中的风险片段在文本中的位置，起始位置从0开始
	Position *[]int32 `json:"position,omitempty"`
}

func (o SegmentResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SegmentResult struct{}"
	}

	return strings.Join([]string{"SegmentResult", string(data)}, " ")
}
