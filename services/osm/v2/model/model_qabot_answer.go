package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QabotAnswer struct {

	// 问题
	Question *string `json:"question,omitempty"`

	// 答案
	Answer *string `json:"answer,omitempty"`

	// 评分
	Score *float64 `json:"score,omitempty"`

	// 主题
	Domain *string `json:"domain,omitempty"`

	// 链接地址
	Link *string `json:"link,omitempty"`

	// 语料id
	QaPairId *string `json:"qa_pair_id,omitempty"`

	// 主题id
	DomainId *string `json:"domain_id,omitempty"`

	// 推荐答案
	TopScoreQuestion *string `json:"top_score_question,omitempty"`

	// 相关问题列表
	RelevanceDetails *[]RelevanceQapair `json:"relevance_details,omitempty"`
}

func (o QabotAnswer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QabotAnswer struct{}"
	}

	return strings.Join([]string{"QabotAnswer", string(data)}, " ")
}
