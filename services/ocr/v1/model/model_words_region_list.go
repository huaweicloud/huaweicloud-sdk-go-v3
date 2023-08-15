package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WordsRegionList 文字区域识别结果列表，输出顺序从左到右，先上后下。
type WordsRegionList struct {

	// 文字识别区域类型。 - text：文本识别区域; - table：表格识别区域。
	Type string `json:"type"`

	// 子区域识别文字块数目。
	WordsBlockCount int32 `json:"words_block_count"`

	// 子区域识别文字块列表，输出顺序从左到右，先上后下。
	WordsBlockList []GeneralTableWordsBlockList `json:"words_block_list"`
}

func (o WordsRegionList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WordsRegionList struct{}"
	}

	return strings.Join([]string{"WordsRegionList", string(data)}, " ")
}
