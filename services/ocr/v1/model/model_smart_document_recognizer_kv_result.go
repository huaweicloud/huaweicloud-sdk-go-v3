package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SmartDocumentRecognizerKvResult struct {

	// 模型识别到的键值对数量。
	KvBlockCount *int32 `json:"kv_block_count,omitempty"`

	// 键值对识别结果列表。
	KvBlockList *[]SmartDocumentRecognizerKvBlock `json:"kv_block_list,omitempty"`
}

func (o SmartDocumentRecognizerKvResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmartDocumentRecognizerKvResult struct{}"
	}

	return strings.Join([]string{"SmartDocumentRecognizerKvResult", string(data)}, " ")
}
