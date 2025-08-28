package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteQuestionAnswerResponse Response Object
type DeleteQuestionAnswerResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteQuestionAnswerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteQuestionAnswerResponse struct{}"
	}

	return strings.Join([]string{"DeleteQuestionAnswerResponse", string(data)}, " ")
}
