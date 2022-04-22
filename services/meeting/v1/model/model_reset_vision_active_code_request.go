package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ResetVisionActiveCodeRequest struct {

	// 请求requestId，用来标识一路请求，用于问题跟踪定位，建议使用uuId，若不携带，则后台自动生成
	XRequestId *string `json:"X-Request-Id,omitempty"`

	// 语言参数，默认为中文zh-CN, 英文为en-US
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// 账号 可通过用户管理中分页查询用户接口获取，对应接口返回userAccount字段。
	Account string `json:"account"`

	Body *ActiveDto `json:"body,omitempty"`
}

func (o ResetVisionActiveCodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetVisionActiveCodeRequest struct{}"
	}

	return strings.Join([]string{"ResetVisionActiveCodeRequest", string(data)}, " ")
}
