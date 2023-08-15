package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DownloadObject struct {

	// 文件名称。
	Name *string `json:"name,omitempty"`

	// 文件大小。
	Size float32 `json:"size,omitempty"`

	// 下载链接。
	DownloadLink *string `json:"download_link,omitempty"`

	// 链接过期时间，格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	LinkExpiredTime *string `json:"link_expired_time,omitempty"`
}

func (o DownloadObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadObject struct{}"
	}

	return strings.Join([]string{"DownloadObject", string(data)}, " ")
}
