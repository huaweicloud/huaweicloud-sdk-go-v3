package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// result字段数据结构说明
type ClassificationResult struct {
	// 待分析文本。

	Content string `json:"content"`
	// 分类标签。 1：广告 0：非广告

	Label int32 `json:"label"`
	// 标签label的置信度。

	Confidence float32 `json:"confidence"`
}

func (o ClassificationResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClassificationResult struct{}"
	}

	return strings.Join([]string{"ClassificationResult", string(data)}, " ")
}
