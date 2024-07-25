package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BackupDownloadLink struct {

	// 文件名
	Name string `json:"name"`

	// 文件大小，单位：KB。
	Size int32 `json:"size"`

	// 文件下载链接。
	DownloadLink string `json:"download_link"`

	// 下载链接过期时间，格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始，Z指时区偏移量，例如北京时间偏移显示为+0800。
	LinkExpiredTime string `json:"link_expired_time"`
}

func (o BackupDownloadLink) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupDownloadLink struct{}"
	}

	return strings.Join([]string{"BackupDownloadLink", string(data)}, " ")
}
