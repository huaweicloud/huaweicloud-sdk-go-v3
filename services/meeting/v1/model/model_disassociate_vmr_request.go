package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DisassociateVmrRequest struct {

	// 请求requestId，用来标识一路请求，用于问题跟踪定位，建议使用uuId，若不携带，则后台自动生成
	XRequestId *string `json:"X-Request-Id,omitempty"`

	// 语言参数，默认为中文zh-CN, 英文为en-US
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// 用户账号
	Account string `json:"account"`

	// 账号类型 * 0：会议账号 * 1：第三方账号。 默认0
	AccountType *int32 `json:"accountType,omitempty"`

	// 云会议室唯一ID列表。 maxLength：100 minLength：1
	Body *[]string `json:"body,omitempty"`
}

func (o DisassociateVmrRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateVmrRequest struct{}"
	}

	return strings.Join([]string{"DisassociateVmrRequest", string(data)}, " ")
}
