package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AnswerInfo struct {

	// 答案
	Answer *string `json:"answer,omitempty"`

	// 链接
	Link *string `json:"link,omitempty"`

	// 问题
	Question *string `json:"question,omitempty"`

	// 语料Id
	QaPairId *string `json:"qa_pair_id,omitempty"`

	// 相关问题详情列表
	RelevanceDetails *[]RelevanceQapair `json:"relevance_details,omitempty"`
}

func (o AnswerInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AnswerInfo struct{}"
	}

	return strings.Join([]string{"AnswerInfo", string(data)}, " ")
}
