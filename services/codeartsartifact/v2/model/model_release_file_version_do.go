package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReleaseFileVersionDo 发布文件版本信息
type ReleaseFileVersionDo struct {

	// 发布库文件的版本
	Version *string `json:"version,omitempty"`

	// 发布库文件的路径
	Path *string `json:"path,omitempty"`

	// 发布库文件的下载链接
	DownloadUrl *string `json:"download_url,omitempty"`
}

func (o ReleaseFileVersionDo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReleaseFileVersionDo struct{}"
	}

	return strings.Join([]string{"ReleaseFileVersionDo", string(data)}, " ")
}
