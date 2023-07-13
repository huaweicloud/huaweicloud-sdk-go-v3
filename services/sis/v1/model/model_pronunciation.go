package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Pronunciation 整体发音打分
type Pronunciation struct {

	// 发音质量综合得分 0-100
	Score float32 `json:"score"`

	// 发音质量好坏 0-100
	Gop float32 `json:"gop"`
}

func (o Pronunciation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Pronunciation struct{}"
	}

	return strings.Join([]string{"Pronunciation", string(data)}, " ")
}
