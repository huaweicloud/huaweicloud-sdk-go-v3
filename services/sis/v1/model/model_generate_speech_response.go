package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GenerateSpeechResponse Response Object
type GenerateSpeechResponse struct {

	// 服务内部的令牌，可用于在日志中追溯具体流程。
	TraceId *string `json:"trace_id,omitempty"`

	Result         *GenerateSpeechRspResult `json:"result,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o GenerateSpeechResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GenerateSpeechResponse struct{}"
	}

	return strings.Join([]string{"GenerateSpeechResponse", string(data)}, " ")
}
