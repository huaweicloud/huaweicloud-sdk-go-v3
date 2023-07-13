package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteCorpAdminsRequest Request Object
type BatchDeleteCorpAdminsRequest struct {

	// 请求requestId，用来标识一路请求，用于问题跟踪定位，建议使用UUID，若不携带，则后台自动生成。
	XRequestId *string `json:"X-Request-Id,omitempty"`

	// 语言参数，默认为中文zh-CN，英文为en-US。
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// 帐号类型。默认0。 * 0：华为云会议帐号。用于帐号/密码鉴权方式 * 1：第三方User ID，用于App ID鉴权方式
	AccountType *int32 `json:"accountType,omitempty"`

	Body *[]string `json:"body,omitempty"`
}

func (o BatchDeleteCorpAdminsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteCorpAdminsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteCorpAdminsRequest", string(data)}, " ")
}
