package model

import (
	"encoding/json"

	"strings"
)

//
type GeneralTableResult struct {
	// 文字区域数目。

	WordsRegionCount int32 `json:"words_region_count"`
	// 文字区域识别结果列表，输出顺序从左到右，先上后下。

	WordsRegionList []WordsRegionList `json:"words_region_list"`
}

func (o GeneralTableResult) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "GeneralTableResult struct{}"
	}

	return strings.Join([]string{"GeneralTableResult", string(data)}, " ")
}
