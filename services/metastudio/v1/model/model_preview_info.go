package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PreviewInfo struct {

	// 文本内容hash值
	TextSha256 *string `json:"text_sha256,omitempty"`

	// 文本对应音频文件下载链接
	AudioFileDownloadUrl *string `json:"audio_file_download_url,omitempty"`
}

func (o PreviewInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PreviewInfo struct{}"
	}

	return strings.Join([]string{"PreviewInfo", string(data)}, " ")
}
