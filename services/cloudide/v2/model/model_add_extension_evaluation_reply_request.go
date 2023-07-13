package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddExtensionEvaluationReplyRequest Request Object
type AddExtensionEvaluationReplyRequest struct {
	Body *EvaluationReply `json:"body,omitempty"`
}

func (o AddExtensionEvaluationReplyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddExtensionEvaluationReplyRequest struct{}"
	}

	return strings.Join([]string{"AddExtensionEvaluationReplyRequest", string(data)}, " ")
}
