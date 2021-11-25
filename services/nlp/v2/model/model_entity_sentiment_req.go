package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type EntitySentimentReq struct {
	// 请求文本。文本编码要求为utf-8，仅支持中文实体情感分析。 限定content+entity长度为512以内，长度超过512时，只检测前512个字符。

	Content string `json:"content"`
	// 请求实体。文本编码要求为utf-8.仅支持中文实体情感分析。 限定content+entity长度为512以内，长度超过512时，只检测前512个字符。

	Entity string `json:"entity"`
	// 取值如下： 3 金融领域

	Type int32 `json:"type"`
}

func (o EntitySentimentReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EntitySentimentReq struct{}"
	}

	return strings.Join([]string{"EntitySentimentReq", string(data)}, " ")
}
