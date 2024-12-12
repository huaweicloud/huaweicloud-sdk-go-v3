package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAsyncTtsJobResponse Response Object
type ShowAsyncTtsJobResponse struct {

	// 音频文件是否已生成完成。该标记为PROCESSING时，应该每隔3秒再次调用本接口获取音频文件(WAITING 等待中,PROCESSING 处理中,SUCCEED 成功,FAILED 失败)
	State *string `json:"state,omitempty"`

	// 音频文件下载链接，有效期为1个小时。
	AudioFileUrl *string `json:"audio_file_url,omitempty"`

	// 字幕文件下载链接，有效期为1个小时。
	AudioInfoFileUrl *string `json:"audio_info_file_url,omitempty"`
	HttpStatusCode   int     `json:"-"`
}

func (o ShowAsyncTtsJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAsyncTtsJobResponse struct{}"
	}

	return strings.Join([]string{"ShowAsyncTtsJobResponse", string(data)}, " ")
}
