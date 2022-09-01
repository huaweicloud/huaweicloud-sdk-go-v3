package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateSessionResponse struct {

	// 会话标识符。
	SessionId *string `json:"session_id,omitempty" xml:"session_id"`

	// 机器人问候语。
	Greeting       *string `json:"greeting,omitempty" xml:"greeting"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSessionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSessionResponse struct{}"
	}

	return strings.Join([]string{"CreateSessionResponse", string(data)}, " ")
}
