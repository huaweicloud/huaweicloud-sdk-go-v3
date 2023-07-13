package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RelevanceQapair struct {

	// 链接
	Link *string `json:"link,omitempty"`

	// 标准问题
	Question *string `json:"question,omitempty"`

	// 标准问题Id
	QaPairId *string `json:"qa_pair_id,omitempty"`
}

func (o RelevanceQapair) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RelevanceQapair struct{}"
	}

	return strings.Join([]string{"RelevanceQapair", string(data)}, " ")
}
