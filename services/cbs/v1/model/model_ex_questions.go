package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExQuestions
type ExQuestions struct {

	// 扩展问题。
	Content string `json:"content"`
}

func (o ExQuestions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExQuestions struct{}"
	}

	return strings.Join([]string{"ExQuestions", string(data)}, " ")
}
