package model

import (
	"encoding/json"

	"strings"
)

//
type HandwritingResult struct {
	// 代表检测识别出来的文字块数目。 \"segment_digit\" 和\"segment\"默认为1。

	WordsBlockCount int32 `json:"words_block_count"`
	// 识别文字块列表，输出顺序从左到右，从上到下。

	WordsBlockList []HandwritingWordsBlockList `json:"words_block_list"`

	ExtractedData *ExtractedData `json:"extracted_data,omitempty"`
}

func (o HandwritingResult) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "HandwritingResult struct{}"
	}

	return strings.Join([]string{"HandwritingResult", string(data)}, " ")
}
