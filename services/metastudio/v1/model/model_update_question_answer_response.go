package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateQuestionAnswerResponse Response Object
type UpdateQuestionAnswerResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateQuestionAnswerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateQuestionAnswerResponse struct{}"
	}

	return strings.Join([]string{"UpdateQuestionAnswerResponse", string(data)}, " ")
}
