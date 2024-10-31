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

	// 返回音频片段结果的策略。可选值如下： false：返回风险等级为非pass的音频片段结果 true：返回所有风险等级的音频片段结果 说明： 1. 默认值为false； 2. 每隔10秒返回一次最近10秒音频流的审核结果。
	ReturnAllResults *bool `json:"return_all_results,omitempty"`
}

func (o AudioStreamCreateRequestData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AudioStreamCreateRequestData struct{}"
	}

	return strings.Join([]string{"AudioStreamCreateRequestData", string(data)}, " ")
}
