package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FailedCustomAnswers struct {

	// 代答库ID
	AnswerId *string `json:"answer_id,omitempty"`

	// 失败原因
	Reason *string `json:"reason,omitempty"`
}

func (o FailedCustomAnswers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FailedCustomAnswers struct{}"
	}

	return strings.Join([]string{"FailedCustomAnswers", string(data)}, " ")
}
