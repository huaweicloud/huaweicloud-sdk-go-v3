package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FileInfo struct {

	// 文件名。
	Name string `json:"name"`

	// 文件大小，单位：KB。
	Size int64 `json:"size"`

	// SQL文件最后一次修改时间，格式为“yyyy-mm-ddThh:mm:ssZ”。  其中，T指某个时间的开始，Z指时区偏移量，例如偏移1个小时显示为+0100。
	UpdatedTime string `json:"updated_time"`

	// 文件下载链接。
	DownloadLink string `json:"download_link"`

	// 下载链接过期时间，格式为“yyyy-mm-ddThh:mm:ssZ”。  其中，T指某个时间的开始，Z指时区偏移量，例如偏移1个小时显示为+0100。
	LinkExpiredTime string `json:"link_expired_time"`
}

func (o FileInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FileInfo struct{}"
	}

	return strings.Join([]string{"FileInfo", string(data)}, " ")
}
