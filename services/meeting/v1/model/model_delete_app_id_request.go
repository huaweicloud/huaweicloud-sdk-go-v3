package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAppIdRequest Request Object
type DeleteAppIdRequest struct {

	// 请求requestId，用来标识一路请求，用于问题跟踪定位，建议使用UUID，若不携带，则后台自动生成。
	XRequestId *string `json:"X-Request-Id,omitempty"`

	// 语言参数，默认为中文zh-CN，英文为en-US。
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// 需要修改的app
	AppId string `json:"app_id"`
}

func (o DeleteAppIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAppIdRequest struct{}"
	}

	return strings.Join([]string{"DeleteAppIdRequest", string(data)}, " ")
}
