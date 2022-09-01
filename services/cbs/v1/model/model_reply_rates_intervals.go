package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type ReplyRatesIntervals struct {

	// 间隔周期开始时间。
	Start string `json:"start" xml:"start"`

	// 间隔周期用户提问总数。
	QuestionCount int64 `json:"question_count" xml:"question_count"`

	// 间隔周期直接回答个数。
	DirectCount int64 `json:"direct_count" xml:"direct_count"`

	// 间隔周期推荐回答个数。
	RecommendCount int64 `json:"recommend_count" xml:"recommend_count"`

	// 间隔周期未匹配个数。
	NotmatchCount int64 `json:"notmatch_count" xml:"notmatch_count"`

	// 间隔周期无效问题个数。
	InvalidCount int64 `json:"invalid_count" xml:"invalid_count"`

	// 间隔周期闲聊匹配个数。
	ChatCount int64 `json:"chat_count" xml:"chat_count"`

	// 间隔周期直接回答比率，保留小数点后三位。
	DirectRate float64 `json:"direct_rate" xml:"direct_rate"`

	// 间隔周期推荐回答比率，保留小数点后三位。
	RecommendRate float64 `json:"recommend_rate" xml:"recommend_rate"`

	// 间隔周期未匹配比率，保留小数点后三位。
	NotmatchRate float64 `json:"notmatch_rate" xml:"notmatch_rate"`

	// 间隔周期闲聊匹配比率，保留小数点后三位。
	ChatRate float64 `json:"chat_rate" xml:"chat_rate"`

	// 间隔周期无效问题比率，保留小数点后三位。
	InvalidRate float64 `json:"invalid_rate" xml:"invalid_rate"`

	// 多轮会话次数。
	DialogCount int64 `json:"dialog_count" xml:"dialog_count"`

	// 多轮会话比例。
	DialogRate float64 `json:"dialog_rate" xml:"dialog_rate"`
}

func (o ReplyRatesIntervals) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReplyRatesIntervals struct{}"
	}

	return strings.Join([]string{"ReplyRatesIntervals", string(data)}, " ")
}
