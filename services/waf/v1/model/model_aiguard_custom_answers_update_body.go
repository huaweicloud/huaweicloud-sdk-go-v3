package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AiguardCustomAnswersUpdateBody struct {

	// 代答库名称
	Name *string `json:"name,omitempty"`

	// 代答库描述
	Description *string `json:"description,omitempty"`

	// 答案列表（全量替换）
	Answers *[]string `json:"answers,omitempty"`
}

func (o AiguardCustomAnswersUpdateBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiguardCustomAnswersUpdateBody struct{}"
	}

	return strings.Join([]string{"AiguardCustomAnswersUpdateBody", string(data)}, " ")
}
