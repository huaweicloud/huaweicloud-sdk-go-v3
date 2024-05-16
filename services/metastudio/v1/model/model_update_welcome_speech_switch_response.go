package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWelcomeSpeechSwitchResponse Response Object
type UpdateWelcomeSpeechSwitchResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateWelcomeSpeechSwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWelcomeSpeechSwitchResponse struct{}"
	}

	return strings.Join([]string{"UpdateWelcomeSpeechSwitchResponse", string(data)}, " ")
}
