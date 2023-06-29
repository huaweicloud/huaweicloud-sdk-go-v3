package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HwCloudSentimentResp result字段数据结构说明
type HwCloudSentimentResp struct {

	// 该文本的分析结果标签，取值如下： 0 负向 1 正向
	Label int32 `json:"label"`

	// 标签label的置信度。小数点精确到（6）位。
	Confidence float32 `json:"confidence"`

	// 待分析文本
	Content string `json:"content"`
}

func (o HwCloudSentimentResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HwCloudSentimentResp struct{}"
	}

	return strings.Join([]string{"HwCloudSentimentResp", string(data)}, " ")
}
