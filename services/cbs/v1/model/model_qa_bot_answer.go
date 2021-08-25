package model

import (
	"encoding/json"

	"strings"
)

//
type QaBotAnswer struct {
	// 问答对ID。

	QaPairId *string `json:"qa_pair_id,omitempty"`
	// 标准问题。

	StQuestion *string `json:"st_question,omitempty"`
	// 知识库答案，包含该字段的回答为直接回答，未包含该字段的是推荐回答。

	Answer *string `json:"answer,omitempty"`
	// 相似度得分，精确到小数点后3位。

	Score *float64 `json:"score,omitempty"`
	// 所属领域。

	Domain *string `json:"domain,omitempty"`
	// 最高评分的扩展问或标准问，当关闭内部阈值处理时返回。

	TopScoreQuestion *string `json:"top_score_question,omitempty"`
}

func (o QaBotAnswer) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "QaBotAnswer struct{}"
	}

	return strings.Join([]string{"QaBotAnswer", string(data)}, " ")
}
