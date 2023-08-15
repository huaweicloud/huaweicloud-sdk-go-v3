package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeFlashAsrResponse Response Object
type RecognizeFlashAsrResponse struct {

	// 服务内部的令牌，可用于在日志中追溯具体调用流程
	TraceId *string `json:"trace_id,omitempty"`

	// 音频时长
	AudioDuration *int32 `json:"audio_duration,omitempty"`

	// 识别结果
	FlashResult    *[]FlashResult `json:"flash_result,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o RecognizeFlashAsrResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeFlashAsrResponse struct{}"
	}

	return strings.Join([]string{"RecognizeFlashAsrResponse", string(data)}, " ")
}
