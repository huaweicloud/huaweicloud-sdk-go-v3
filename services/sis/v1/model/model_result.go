package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Result struct {

	// 调用成功表示识别出的内容。
	Text string `json:"text"`

	// 调用成功表示识别出的置信度，取值范围：0~1。
	Score float32 `json:"score"`

	// 分词信息列表
	WordInfo *[]WordInfo `json:"word_info,omitempty"`
}

func (o Result) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Result struct{}"
	}

	return strings.Join([]string{"Result", string(data)}, " ")
}
