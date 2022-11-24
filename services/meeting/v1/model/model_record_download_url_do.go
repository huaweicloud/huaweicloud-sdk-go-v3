package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 单个录制文件下载链接信息。
type RecordDownloadUrlDo struct {

	// 下载鉴权token，下载文件时，使用该token鉴权。（一小时内有效，使用后立即失效）。
	Token *string `json:"token,omitempty"`

	// 文件类型。 * Aux：辅流（会议中的共享画面；分辨率为720p） * Hd：高清（会议中的视频画面；分辨率和会议中视频画面的分辨率一致，1080p或者720p） * Sd：标清（会议中视频画面和共享画面的合成画面，视频画面是大画面，共享画面是小画面，共享画面布局在右下方；分辨率为4CIF） > 单个MP4文件大小不超过1GB。
	FileType *string `json:"fileType,omitempty"`

	// 文件下载url，最大1000个字符。
	Url *string `json:"url,omitempty"`
}

func (o RecordDownloadUrlDo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecordDownloadUrlDo struct{}"
	}

	return strings.Join([]string{"RecordDownloadUrlDo", string(data)}, " ")
}
