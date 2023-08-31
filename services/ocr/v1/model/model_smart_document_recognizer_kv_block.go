package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SmartDocumentRecognizerKvBlock struct {

	// key-value对（键值对）中的key，例如“姓名：小明”中的“姓名”
	Key *string `json:"key,omitempty"`

	// key-value对（键值对）中的value，例如“姓名：小明”中的“小明”
	Value *string `json:"value,omitempty"`

	// 该键值对中所包含的文本框数量。
	WordsBlockCount *int32 `json:"words_block_count,omitempty"`

	// 文本框识别结果列表。
	WordsBlockList *[]SmartDocumentRecognizerKvWordsBlock `json:"words_block_list,omitempty"`
}

func (o SmartDocumentRecognizerKvBlock) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmartDocumentRecognizerKvBlock struct{}"
	}

	return strings.Join([]string{"SmartDocumentRecognizerKvBlock", string(data)}, " ")
}
