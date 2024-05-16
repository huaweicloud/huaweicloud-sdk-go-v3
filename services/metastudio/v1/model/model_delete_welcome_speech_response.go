package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteWelcomeSpeechResponse Response Object
type DeleteWelcomeSpeechResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteWelcomeSpeechResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWelcomeSpeechResponse struct{}"
	}

	return strings.Join([]string{"DeleteWelcomeSpeechResponse", string(data)}, " ")
}
