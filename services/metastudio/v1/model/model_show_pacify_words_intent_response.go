package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPacifyWordsIntentResponse Response Object
type ShowPacifyWordsIntentResponse struct {

	// 安抚话术意图信息
	Data *[]PacifyWordsIntentInfo `json:"data,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPacifyWordsIntentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPacifyWordsIntentResponse struct{}"
	}

	return strings.Join([]string{"ShowPacifyWordsIntentResponse", string(data)}, " ")
}
