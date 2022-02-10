package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 整体流利度打分
type Fluency struct {
	// 流利度综合得分 0-100

	Score float32 `json:"score"`
	// 韵律得分 0-100 韵律指音素在单词和句子中的发音长度是否得当

	Rhythm float32 `json:"rhythm"`
	// 连贯性得分 0-100

	Cohesion float32 `json:"cohesion"`
}

func (o Fluency) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Fluency struct{}"
	}

	return strings.Join([]string{"Fluency", string(data)}, " ")
}
