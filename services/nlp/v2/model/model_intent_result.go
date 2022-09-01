package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type IntentResult struct {

	// 标签label的置信度。
	Confidence float32 `json:"confidence" xml:"confidence"`

	// 文本的意图标签。标签共有以下9类，取值如下： weather：天气 time：报时 news：新闻 joke：笑话 translation：翻译 notification：提醒 alarm：闹钟 music：音乐 other：其它
	Label string `json:"label" xml:"label"`

	// slot数据结构
	Slots []Slot `json:"slots" xml:"slots"`

	// 返回待分析文本。
	Text string `json:"text" xml:"text"`
}

func (o IntentResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IntentResult struct{}"
	}

	return strings.Join([]string{"IntentResult", string(data)}, " ")
}
