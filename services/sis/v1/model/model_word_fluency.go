package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WordFluency 单词的流利度打分
type WordFluency struct {

	//
	Score float32 `json:"score"`

	//
	Rhythm float32 `json:"rhythm"`
}

func (o WordFluency) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WordFluency struct{}"
	}

	return strings.Join([]string{"WordFluency", string(data)}, " ")
}
