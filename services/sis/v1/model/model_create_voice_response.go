package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVoiceResponse Response Object
type CreateVoiceResponse struct {

	// 服务内部的令牌，可用于在日志中追溯具体流程。
	TraceId *string `json:"trace_id,omitempty"`

	Result         *RegisterVoiceResponseBodyResult `json:"result,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o CreateVoiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVoiceResponse struct{}"
	}

	return strings.Join([]string{"CreateVoiceResponse", string(data)}, " ")
}
