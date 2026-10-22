package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AiguardCustomAnswersBody struct {

	// 代答库名称，同project唯一
	Name string `json:"name"`

	// 代答库描述
	Description *string `json:"description,omitempty"`

	// 答案列表，至少1条
	Answers []string `json:"answers"`
}

func (o AiguardCustomAnswersBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiguardCustomAnswersBody struct{}"
	}

	return strings.Join([]string{"AiguardCustomAnswersBody", string(data)}, " ")
}
