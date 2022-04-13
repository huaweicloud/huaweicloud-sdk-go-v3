package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type WordInfo struct {
	// 起始时间

	StartTime *int32 `json:"start_time,omitempty"`
	// 结束时间

	EndTime *int32 `json:"end_time,omitempty"`
	// 分词

	Word *string `json:"word,omitempty"`
}

func (o WordInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WordInfo struct{}"
	}

	return strings.Join([]string{"WordInfo", string(data)}, " ")
}
