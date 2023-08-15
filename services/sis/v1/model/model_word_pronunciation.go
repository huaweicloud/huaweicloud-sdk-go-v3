package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WordPronunciation  单词发音打分
type WordPronunciation struct {

	// 单词发音综合得分 0-100
	Score float32 `json:"score"`

	// 单词发音好坏得分 0-100
	Gop float32 `json:"gop"`
}

func (o WordPronunciation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WordPronunciation struct{}"
	}

	return strings.Join([]string{"WordPronunciation", string(data)}, " ")
}
