package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TtsCallBackConfig struct {

	// 回调URL。 回调请求body为json格式，带参数如下： status: FINISHED或ERROR或者WAITING job_id: 任务id audio_file_download_url: 音频文件路径 subtitle_file_download_url: 字幕文件路径 audio_duration: 音频时长（秒）
	CallbackUrl string `json:"callback_url"`
}

func (o TtsCallBackConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TtsCallBackConfig struct{}"
	}

	return strings.Join([]string{"TtsCallBackConfig", string(data)}, " ")
}
