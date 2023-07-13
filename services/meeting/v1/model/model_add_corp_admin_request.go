package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddCorpAdminRequest Request Object
type AddCorpAdminRequest struct {

	// 请求requestId，用来标识一路请求，用于问题跟踪定位，建议使用UUID，若不携带，则后台自动生成。
	XRequestId *string `json:"X-Request-Id,omitempty"`

	// 语言参数，默认为中文zh-CN，英文为en-US。
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// 帐号类型。默认0。 * 0：会议帐号 * 1：表示第三方帐号。
	AccountType *int32 `json:"accountType,omitempty"`

	Body *CorpAdminDto `json:"body,omitempty"`
}

func (o AddCorpAdminRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddCorpAdminRequest struct{}"
	}

	return strings.Join([]string{"AddCorpAdminRequest", string(data)}, " ")
}
