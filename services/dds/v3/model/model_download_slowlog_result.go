package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DownloadSlowlogResult struct {
	// 节点名称。

	NodeName string `json:"node_name"`
	// 生成的下载文件名。

	FileName string `json:"file_name"`
	// 当前链接的生成状态。 - SUCCESS，表示下载链接已经生成完成。 - EXPORTING，表示正在生成文件，准备下载链接。 - FAILED，表示存在日志文件准备失败。

	Status string `json:"status"`
	// 文件大小，单位为 KB。

	FileSize string `json:"file_size"`
	// 下载链接。注意：下载链接在更新时间的 15分钟内有效，超出时间会重新获取。

	FileLink string `json:"file_link"`
	// 更新时间。

	UpdateAt int64 `json:"update_at"`
}

func (o DownloadSlowlogResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadSlowlogResult struct{}"
	}

	return strings.Join([]string{"DownloadSlowlogResult", string(data)}, " ")
}
