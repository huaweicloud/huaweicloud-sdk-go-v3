package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FlashScoreResult struct {

	// 调用成功表示识别出的内容。
	Text string `json:"text" xml:"text"`

	// 调用成功表示识别出的置信度，取值范围：0~1。
	Score float64 `json:"score" xml:"score"`

	// 分词信息列表
	WordInfo *[]WordInfo `json:"word_info,omitempty" xml:"word_info"`
}

func (o FlashScoreResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlashScoreResult struct{}"
	}

	return strings.Join([]string{"FlashScoreResult", string(data)}, " ")
}
