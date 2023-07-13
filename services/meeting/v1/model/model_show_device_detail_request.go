package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDeviceDetailRequest Request Object
type ShowDeviceDetailRequest struct {

	// 请求requestId，用来标识一路请求，用于问题跟踪定位，建议使用UUID，若不携带，则后台自动生成。
	XRequestId *string `json:"X-Request-Id,omitempty"`

	// 语言参数，默认为中文zh-CN，英文为en-US。
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// 终端SN号，仅可包含数字、字母和下划线。
	Sn string `json:"sn"`
}

func (o ShowDeviceDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeviceDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowDeviceDetailRequest", string(data)}, " ")
}
