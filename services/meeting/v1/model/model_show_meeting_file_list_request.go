package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowMeetingFileListRequest struct {
	// 请求requestId，用来标识一路请求，用于问题跟踪定位，建议使用uuId，若不携带，则后台自动生成

	XRequestId *string `json:"X-Request-Id,omitempty"`
	// 语言参数，默认为中文zh-CN, 英文为en-US

	AcceptLanguage *string `json:"Accept-Language,omitempty"`
	// 二维码内容中所有key/value存为的map集合

	Body map[string]string `json:"body,omitempty"`
}

func (o ShowMeetingFileListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMeetingFileListRequest struct{}"
	}

	return strings.Join([]string{"ShowMeetingFileListRequest", string(data)}, " ")
}
