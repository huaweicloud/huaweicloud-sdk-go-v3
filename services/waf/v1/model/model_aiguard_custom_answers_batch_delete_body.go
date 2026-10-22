package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AiguardCustomAnswersBatchDeleteBody struct {

	// 代答库ID列表
	AnswerIds []string `json:"answer_ids"`
}

func (o AiguardCustomAnswersBatchDeleteBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiguardCustomAnswersBatchDeleteBody struct{}"
	}

	return strings.Join([]string{"AiguardCustomAnswersBatchDeleteBody", string(data)}, " ")
}
