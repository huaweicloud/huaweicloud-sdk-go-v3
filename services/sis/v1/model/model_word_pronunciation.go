package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//  单词发音打分
type WordPronunciation struct {

	// 单词发音综合得分 0-100
	Score float32 `json:"score" xml:"score"`

	// 单词发音好坏得分 0-100
	Gop float32 `json:"gop" xml:"gop"`
}

func (o WordPronunciation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WordPronunciation struct{}"
	}

	return strings.Join([]string{"WordPronunciation", string(data)}, " ")
}
