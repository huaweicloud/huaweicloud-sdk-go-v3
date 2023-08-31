package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SmartDocumentRecognizerOcrResult struct {

	// 图片朝向
	Direction *float32 `json:"direction,omitempty"`

	// 识别文字块数目。
	WordsBlockCount *int32 `json:"words_block_count,omitempty"`

	// 识别文字块列表，输出顺序从左到右，先上后下。
	WordsBlockList *[]SmartDocumentRecognizerWordsBlockList `json:"words_block_list,omitempty"`
}

func (o SmartDocumentRecognizerOcrResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmartDocumentRecognizerOcrResult struct{}"
	}

	return strings.Join([]string{"SmartDocumentRecognizerOcrResult", string(data)}, " ")
}
