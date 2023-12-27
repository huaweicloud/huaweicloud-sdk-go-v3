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
}

func (o CreateFasReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFasReq struct{}"
	}

	return strings.Join([]string{"CreateFasReq", string(data)}, " ")
}
