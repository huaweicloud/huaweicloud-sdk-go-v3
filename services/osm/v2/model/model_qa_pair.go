package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QaPair struct {

	// 主题
	Domain *string `json:"domain,omitempty"`

	// 链接
	Link *string `json:"link,omitempty"`

	// 问题
	Question *string `json:"question,omitempty"`

	// 答案列表
	Answers *[]Answer `json:"answers,omitempty"`

	// 语料Id
	QaPairId *string `json:"qa_pair_id,omitempty"`

	// 扩展问题列表
	ExQuestions *[]ExtendQuestion `json:"ex_questions,omitempty"`

	// 相关问题列表
	RelatedQuestionIds *[]string `json:"related_question_ids,omitempty"`
}

func (o QaPair) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QaPair struct{}"
	}

	return strings.Join([]string{"QaPair", string(data)}, " ")
}
