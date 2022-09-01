package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 单个录制文件下载链接信息
type RecordDownloadUrlDo struct {

	// 下载鉴权token
	Token *string `json:"token,omitempty" xml:"token"`

	// 文件类型
	FileType *string `json:"fileType,omitempty" xml:"fileType"`

	// 录制文件下载URL
	Url *string `json:"url,omitempty" xml:"url"`
}

func (o RecordDownloadUrlDo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecordDownloadUrlDo struct{}"
	}

	return strings.Join([]string{"RecordDownloadUrlDo", string(data)}, " ")
}
