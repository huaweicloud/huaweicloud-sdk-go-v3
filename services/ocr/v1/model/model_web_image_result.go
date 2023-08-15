package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WebImageResult
type WebImageResult struct {

	// 代表检测识别出来的文字块数目。
	WordsBlockCount int32 `json:"words_block_count"`

	// 识别文字块列表，输出顺序从左到右，从上到下。
	WordsBlockList []WebImageWordsBlockList `json:"words_block_list"`

	ExtractedData *WebImageExtractedData `json:"extracted_data"`
}

func (o WebImageResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebImageResult struct{}"
	}

	return strings.Join([]string{"WebImageResult", string(data)}, " ")
}
