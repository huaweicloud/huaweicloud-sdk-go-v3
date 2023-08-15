package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetActivecodeRequest Request Object
type ResetActivecodeRequest struct {

	// 请求requestId，用来标识一路请求，用于问题跟踪定位，建议使用UUID，若不携带，则后台自动生成。
	XRequestId *string `json:"X-Request-Id,omitempty"`

	// 语言参数，默认为中文zh-CN，英文为en-US。
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// 终端序列号，仅可包含数字、字母和下划线。
	Sn string `json:"sn"`

	Body *ActiveDto `json:"body,omitempty"`
}

func (o ResetActivecodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetActivecodeRequest struct{}"
	}

	return strings.Join([]string{"ResetActivecodeRequest", string(data)}, " ")
}
