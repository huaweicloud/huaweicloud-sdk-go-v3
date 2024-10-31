package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CaptureResultUrlVo struct {

	// 下载链接提取码，用于打开下载链接时使用。
	Captcha *string `json:"captcha,omitempty"`

	// 下载链接过期时间
	Expires *int64 `json:"expires,omitempty"`

	// 抓包文件列表，当环境不支持obs文件夹分享时使用。当此字段存在时，无captch，expires，url返回值。
	FileList *[]CaptureFile `json:"file_list,omitempty"`

	RequestHeader *HostHeaderInfo `json:"request_header,omitempty"`

	// 下载链接
	Url *string `json:"url,omitempty"`
}

func (o CaptureResultUrlVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CaptureResultUrlVo struct{}"
	}

	return strings.Join([]string{"CaptureResultUrlVo", string(data)}, " ")
}
