package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchDeleteProgramsRequest struct {

	// 请求requestId，用来标识一路请求，用于问题跟踪定位，建议使用UUID，若不携带，则后台自动生成。
	XRequestId *string `json:"X-Request-Id,omitempty"`

	// 语言参数，默认为中文zh-CN，英文为en-US。
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	Body *[]string `json:"body,omitempty"`
}

func (o BatchDeleteProgramsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteProgramsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteProgramsRequest", string(data)}, " ")
}
