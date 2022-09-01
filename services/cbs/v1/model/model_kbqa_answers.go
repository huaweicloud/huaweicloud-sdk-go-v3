package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type KbqaAnswers struct {

	// 答案。
	Answer string `json:"answer" xml:"answer"`

	// 答案评分。
	Score float64 `json:"score" xml:"score"`
}

func (o KbqaAnswers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KbqaAnswers struct{}"
	}

	return strings.Join([]string{"KbqaAnswers", string(data)}, " ")
}
