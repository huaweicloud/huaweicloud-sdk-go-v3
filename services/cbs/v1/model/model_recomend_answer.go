package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type RecomendAnswer struct {
	// 问答对ID。

	QaPairId *string `json:"qa_pair_id,omitempty"`
	// 标准问题。

	StQuestion *string `json:"st_question,omitempty"`
	// 相似度得分，精确到小数点后3位。

	Score *float64 `json:"score,omitempty"`
	// 所属领域。

	Domain *string `json:"domain,omitempty"`
	// 最高评分的扩展问或标准问。

	TopScoreQuestion *string `json:"top_score_question,omitempty"`
}

func (o RecomendAnswer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecomendAnswer struct{}"
	}

	return strings.Join([]string{"RecomendAnswer", string(data)}, " ")
}
