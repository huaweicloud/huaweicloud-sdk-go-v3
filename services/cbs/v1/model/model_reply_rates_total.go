package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type ReplyRatesTotal struct {
	// 用户提问总数。

	QuestionCount int64 `json:"question_count"`
	// 直接回答个数。

	DirectCount int64 `json:"direct_count"`
	// 推荐回答个数。

	RecommendCount int64 `json:"recommend_count"`
	// 未匹配个数。

	NotmatchCount int64 `json:"notmatch_count"`
	// 直接回答比率，保留小数点后三位。

	DirectRate float64 `json:"direct_rate"`
	// 推荐回答比率，保留小数点后三位。

	RecommendRate float64 `json:"recommend_rate"`
	// 未匹配比率，保留小数点后三位。

	NotmatchRate float64 `json:"notmatch_rate"`
	// 多轮对话次数。

	DialogCount int64 `json:"dialog_count"`
	// 多轮对话比例。

	DialogRate float64 `json:"dialog_rate"`
	// 无效问题次数。

	InvalidCount int64 `json:"invalid_count"`
	// 无效问题比例。

	InvalidRate float64 `json:"invalid_rate"`
	// 闲聊匹配次数。

	ChatCount int64 `json:"chat_count"`
	// 闲聊比例。

	ChatRate float64 `json:"chat_rate"`
}

func (o ReplyRatesTotal) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReplyRatesTotal struct{}"
	}

	return strings.Join([]string{"ReplyRatesTotal", string(data)}, " ")
}
