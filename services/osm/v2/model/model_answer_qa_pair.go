package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AnswerQaPair struct {

	// 相似度得分
	Score *float64 `json:"score,omitempty"`

	// 答案
	Answer *string `json:"answer,omitempty"`

	// 主题
	Domain *string `json:"domain,omitempty"`

	// 链接
	Link *string `json:"link,omitempty"`

	// 问题
	Question *string `json:"question,omitempty"`

	// 语料Id
	QaPairId *string `json:"qa_pair_id,omitempty"`

	// 相关问题详情列表
	RelevanceDetails *[]RelevanceQapair `json:"relevance_details,omitempty"`

	// 技能意图识别是否结束
	TaskComplete *bool `json:"task_complete,omitempty"`
}

func (o AnswerQaPair) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AnswerQaPair struct{}"
	}

	return strings.Join([]string{"AnswerQaPair", string(data)}, " ")
}
