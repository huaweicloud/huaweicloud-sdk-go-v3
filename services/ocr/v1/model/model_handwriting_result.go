package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HandwritingResult
type HandwritingResult struct {

	// 代表检测识别出来的文字块数目。 \"segment_digit\" 和\"segment\"默认为1。
	WordsBlockCount int32 `json:"words_block_count"`

	// 识别文字块列表，输出顺序从左到右，从上到下。
	WordsBlockList []HandwritingWordsBlockList `json:"words_block_list"`
}

func (o HandwritingResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HandwritingResult struct{}"
	}

	return strings.Join([]string{"HandwritingResult", string(data)}, " ")
}
