package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EvaluationReply 评论回复,其中evaluation_id和reply_id之中必须选择一个填写
type EvaluationReply struct {

	// 所在评论id
	EvaluationId *string `json:"evaluation_id,omitempty"`

	// 回复评论的id 空表示回复主评论
	ReplyId *string `json:"reply_id,omitempty"`

	// 评论或回复内容
	Text string `json:"text"`
}

func (o EvaluationReply) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EvaluationReply struct{}"
	}

	return strings.Join([]string{"EvaluationReply", string(data)}, " ")
}
