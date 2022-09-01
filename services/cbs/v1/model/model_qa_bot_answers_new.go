package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type QaBotAnswersNew struct {

	// 问答机器人回复。
	Answers *[]QaBotAnswer `json:"answers,omitempty" xml:"answers"`

	// 问答机器人推荐问题
	RecommendAnswers *[]RecomendAnswer `json:"recommend_answers,omitempty" xml:"recommend_answers"`
}

func (o QaBotAnswersNew) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QaBotAnswersNew struct{}"
	}

	return strings.Join([]string{"QaBotAnswersNew", string(data)}, " ")
}
