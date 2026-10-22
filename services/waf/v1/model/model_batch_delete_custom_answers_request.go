package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteCustomAnswersRequest Request Object
type BatchDeleteCustomAnswersRequest struct {
	Body *AiguardCustomAnswersBatchDeleteBody `json:"body,omitempty"`
}

func (o BatchDeleteCustomAnswersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteCustomAnswersRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteCustomAnswersRequest", string(data)}, " ")
}
