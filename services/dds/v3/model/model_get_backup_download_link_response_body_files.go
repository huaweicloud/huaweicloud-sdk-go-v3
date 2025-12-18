package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GetBackupDownloadLinkResponseBodyFiles struct {

	// 文件名。
	Name string `json:"name"`

	// 文件大小，单位为KB。
	Size int64 `json:"size"`

	// 文件下载链接。
	DownloadLink string `json:"download_link"`

	// 下载链接过期时间，格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始，Z指时区偏移量，例如UTC时间偏移显示为+0000。
	LinkExpiredTime string `json:"link_expired_time"`

	// 组ID。
	GroupId *string `json:"group_id,omitempty"`

	// 组名。
	GroupName *string `json:"group_name,omitempty"`
}

func (o GetBackupDownloadLinkResponseBodyFiles) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetBackupDownloadLinkResponseBodyFiles struct{}"
	}

	return strings.Join([]string{"GetBackupDownloadLinkResponseBodyFiles", string(data)}, " ")
}
