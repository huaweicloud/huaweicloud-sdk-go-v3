package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Word 单个单词的发音评测结果
type Word struct {

	// 去除了所有标点符号后的原始文本 前端调用接口后推荐使用“​text​”来在UI 中展示结果
	Text string `json:"text"`

	// 接口接收的原始文本
	TextOriginal string `json:"text_original"`

	// 原始文本规范化后切分成的单词 如175 会 规范为 [\"one\", \"\"hundred\", \"and\", \"seventy\", \"five\"]
	TextNormalised []string `json:"text_normalised"`

	// 是否命中模型发音字典 如果未命中，则表明会根据发音规律推测正确发音
	OutOfVocabulary *bool `json:"out_of_vocabulary,omitempty"`

	// 起始时间
	StartTime *float32 `json:"start_time,omitempty"`

	// 结束时间
	EndTime *float32 `json:"end_time,omitempty"`

	// 综合评分
	Score *float32 `json:"score,omitempty"`

	Pronunciation *WordPronunciation `json:"pronunciation,omitempty"`

	Fluency *WordFluency `json:"fluency,omitempty"`

	// 音节打分表
	Phonemes *[]Phoneme `json:"phonemes,omitempty"`
}

func (o Word) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Word struct{}"
	}

	return strings.Join([]string{"Word", string(data)}, " ")
}
