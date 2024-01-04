package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateFasReq struct {

	// 语音驱动音频文件下载URL，格式为AAC或者MP3
	AudioFileDownloadUrl string `json:"audio_file_download_url"`

	// 期望的输出帧率
	FrameRate int32 `json:"frame_rate"`

	// 情绪： 0：平静（默认） 1：开心 2：哀伤 3：愤怒
	Emotion *int32 `json:"emotion,omitempty"`
}

func (o CreateFasReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFasReq struct{}"
	}

	return strings.Join([]string{"CreateFasReq", string(data)}, " ")
}
