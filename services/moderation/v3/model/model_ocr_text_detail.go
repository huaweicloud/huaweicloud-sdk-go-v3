package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OcrTextDetail struct {

	// 命中的风险片段
	Segment *string `json:"segment,omitempty"`

	// 命中的自定义词库名称。 命中自定义词库时，才会返回当前字段。
	GlossaryName *string `json:"glossary_name,omitempty"`
}

func (o OcrTextDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OcrTextDetail struct{}"
	}

	return strings.Join([]string{"OcrTextDetail", string(data)}, " ")
}
