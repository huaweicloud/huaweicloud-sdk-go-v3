package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EvaluationAccusation 其中evaluation_id和reply_id之中必须选择一个填写
type EvaluationAccusation struct {

	// 举报内容
	Content string `json:"content"`

	// 评论id
	EvaluationId *string `json:"evaluation_id,omitempty"`

	// 回复id
	ReplyId *string `json:"reply_id,omitempty"`
}

func (o EvaluationAccusation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EvaluationAccusation struct{}"
	}

	return strings.Join([]string{"EvaluationAccusation", string(data)}, " ")
}
