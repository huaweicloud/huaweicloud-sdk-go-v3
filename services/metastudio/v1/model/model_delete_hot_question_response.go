package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHotQuestionResponse Response Object
type DeleteHotQuestionResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteHotQuestionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHotQuestionResponse struct{}"
	}

	return strings.Join([]string{"DeleteHotQuestionResponse", string(data)}, " ")
}
