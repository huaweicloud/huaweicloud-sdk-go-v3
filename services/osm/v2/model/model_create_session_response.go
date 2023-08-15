package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSessionResponse Response Object
type CreateSessionResponse struct {

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 会话Id
	SessionId      *string `json:"session_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSessionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSessionResponse struct{}"
	}

	return strings.Join([]string{"CreateSessionResponse", string(data)}, " ")
}
