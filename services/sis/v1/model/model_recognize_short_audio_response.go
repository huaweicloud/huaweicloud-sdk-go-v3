package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeShortAudioResponse Response Object
type RecognizeShortAudioResponse struct {

	// 服务内部的令牌，可用于在日志中追溯具体流程，调用失败无此字段。  在某些错误情况下可能没有此令牌字符串。
	TraceId *string `json:"trace_id,omitempty"`

	Result         *Result `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RecognizeShortAudioResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeShortAudioResponse struct{}"
	}

	return strings.Join([]string{"RecognizeShortAudioResponse", string(data)}, " ")
}
