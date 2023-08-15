package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SimpleQaPair struct {

	// 问题
	Question *string `json:"question,omitempty"`

	// 主题
	Domain *string `json:"domain,omitempty"`

	// 语料Id
	QaPairId *string `json:"qa_pair_id,omitempty"`
}

func (o SimpleQaPair) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SimpleQaPair struct{}"
	}

	return strings.Join([]string{"SimpleQaPair", string(data)}, " ")
}
