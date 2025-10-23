package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GeneralTextResult
type GeneralTextResult struct {

	// 图片朝向，仅当detect_direction为true时，该字段有效。返回图片逆时针旋转角度，值区间为[0， 359],保留四位小数。 当detect_direction为false时，该字段值为 -1。
	Direction float32 `json:"direction"`

	// 识别文字块数目。
	WordsBlockCount int32 `json:"words_block_count"`

	// 识别文字块列表，输出顺序从左到右，先上后下。
	WordsBlockList []GeneralTextWordsBlockList `json:"words_block_list"`

	// 所有文字块拼接的识别结果，同一行的文字块使用“\\t”拼接，不同行的文字块使用“\\n”拼接。 当return_markdown_result为true时，返回该字段值，否则，不返回该字段。
	MarkdownResult *string `json:"markdown_result,omitempty"`
}

func (o GeneralTextResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GeneralTextResult struct{}"
	}

	return strings.Join([]string{"GeneralTextResult", string(data)}, " ")
}
