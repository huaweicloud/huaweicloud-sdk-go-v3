package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type WordInfo struct {

	// 起始时间
	StartTime *int32 `json:"start_time,omitempty" xml:"start_time"`

	// 结束时间
	EndTime *int32 `json:"end_time,omitempty" xml:"end_time"`

	// 分词
	Word *string `json:"word,omitempty" xml:"word"`
}

func (o WordInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WordInfo struct{}"
	}

	return strings.Join([]string{"WordInfo", string(data)}, " ")
}
