package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AudioStreamCreateRequestData 音频流数据输入
type AudioStreamCreateRequestData struct {

	// 音频流url地址，支持rtmp、rtmps、hls、http、https等主流协议。
	Url string `json:"url"`

	// 指定音频流中语种类型 zh: 中文,默认值为zh
	Language *string `json:"language,omitempty"`
}

func (o AudioStreamCreateRequestData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AudioStreamCreateRequestData struct{}"
	}

	return strings.Join([]string{"AudioStreamCreateRequestData", string(data)}, " ")
}
