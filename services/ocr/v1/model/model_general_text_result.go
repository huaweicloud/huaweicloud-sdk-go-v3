package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type GeneralTextResult struct {
	// 图片朝向，仅当detect_direction为true时，该字段有效。返回图片逆时针旋转角度，值区间为[0， 359]。当detect_direction为false时，该字段值为 -1。

	Direction int32 `json:"direction"`
	// 识别文字块数目。

	WordsBlockCount int32 `json:"words_block_count"`
	// 识别文字块列表，输出顺序从左到右，先上后下。

	WordsBlockList []GeneralTextWordsBlockList `json:"words_block_list"`
}

func (o GeneralTextResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GeneralTextResult struct{}"
	}

	return strings.Join([]string{"GeneralTextResult", string(data)}, " ")
}
