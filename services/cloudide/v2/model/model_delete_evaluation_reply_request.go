package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEvaluationReplyRequest Request Object
type DeleteEvaluationReplyRequest struct {

	// 回复id
	ReplyId int64 `json:"reply_id"`
}

func (o DeleteEvaluationReplyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEvaluationReplyRequest struct{}"
	}

	return strings.Join([]string{"DeleteEvaluationReplyRequest", string(data)}, " ")
}
