package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type RecomendAnswer struct {

	// 问答对ID。
	QaPairId *string `json:"qa_pair_id,omitempty" xml:"qa_pair_id"`

	// 标准问题。
	StQuestion *string `json:"st_question,omitempty" xml:"st_question"`

	// 相似度得分，精确到小数点后3位。
	Score *float64 `json:"score,omitempty" xml:"score"`

	// 所属领域。
	Domain *string `json:"domain,omitempty" xml:"domain"`

	// 最高评分的扩展问或标准问。
	TopScoreQuestion *string `json:"top_score_question,omitempty" xml:"top_score_question"`
}

func (o RecomendAnswer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecomendAnswer struct{}"
	}

	return strings.Join([]string{"RecomendAnswer", string(data)}, " ")
}
