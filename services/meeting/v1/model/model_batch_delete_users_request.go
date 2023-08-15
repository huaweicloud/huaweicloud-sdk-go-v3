package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteUsersRequest Request Object
type BatchDeleteUsersRequest struct {

	// 请求requestId，用来标识一路请求，用于问题跟踪定位，建议使用UUID，若不携带，则后台自动生成。
	XRequestId *string `json:"X-Request-Id,omitempty"`

	// 语言参数，默认为中文zh-CN，英文为en-US。
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// 帐号类型。默认0。 * 0：会议帐号 * 1：第三方帐号
	AccountType *int32 `json:"accountType,omitempty"`

	Body *[]string `json:"body,omitempty"`
}

func (o BatchDeleteUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteUsersRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteUsersRequest", string(data)}, " ")
}
