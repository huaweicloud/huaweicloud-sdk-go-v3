package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDeviceStatusRequest struct {

	// 请求requestId，用来标识一路请求，用于问题跟踪定位，建议使用uuId，若不携带，则后台自动生成
	XRequestID *string `json:"X-Request-ID,omitempty" xml:"X-Request-ID"`

	// 语言参数，默认为中文zh-CN, 英文为en-US
	AcceptLanguage *string `json:"Accept-Language,omitempty" xml:"Accept-Language"`

	// 所请求查询的终端的号码。
	Body *[]string `json:"body,omitempty" xml:"body"`
}

func (o ShowDeviceStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeviceStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowDeviceStatusRequest", string(data)}, " ")
}
