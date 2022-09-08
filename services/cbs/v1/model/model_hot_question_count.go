package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type HotQuestionCount struct {

	// 问答对ID。
	QaPairId *string `json:"qa_pair_id,omitempty"`

	// 标准问题。
	StQuestion *string `json:"st_question,omitempty"`

	// 标准问题所属领域。
	Domain *string `json:"domain,omitempty"`
}

func (o HotQuestionCount) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HotQuestionCount struct{}"
	}

	return strings.Join([]string{"HotQuestionCount", string(data)}, " ")
}
