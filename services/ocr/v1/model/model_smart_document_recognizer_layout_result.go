package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SmartDocumentRecognizerLayoutResult struct {

	// 模型识别到的文档版面区域数量。
	LayoutBlockCount *int32 `json:"layout_block_count,omitempty"`

	// 文档版面区域识别结果列表。
	LayoutBlockList *[]SmartDocumentRecognizerLayoutBlock `json:"layout_block_list,omitempty"`
}

func (o SmartDocumentRecognizerLayoutResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmartDocumentRecognizerLayoutResult struct{}"
	}

	return strings.Join([]string{"SmartDocumentRecognizerLayoutResult", string(data)}, " ")
}
