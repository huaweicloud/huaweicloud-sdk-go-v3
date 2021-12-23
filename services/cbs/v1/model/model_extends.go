package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type Extends struct {
	TagIds *Tag `json:"tag_ids,omitempty"`
	// 领域列表，多个领域用分号隔开。如果设置了领域且领域不为空，就从这些领域中匹配答案，否则就从该用户的全部知识库匹配答案。  当前最多支持10个领域。

	DomainIds *[]string `json:"domain_ids,omitempty"`
	// 问题来源 其他支持用户自定义，最终体现在问答日志里

	Source *string `json:"source,omitempty"`
	// 是否返回所有类型的回答

	ReturnAllAnswers *bool `json:"return_all_answers,omitempty"`
}

func (o Extends) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Extends struct{}"
	}

	return strings.Join([]string{"Extends", string(data)}, " ")
}
