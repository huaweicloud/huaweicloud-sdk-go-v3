package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchCorpAdminsRequest Request Object
type SearchCorpAdminsRequest struct {

	// 请求requestId，用来标识一路请求，用于问题跟踪定位，建议使用UUID，若不携带，则后台自动生成。
	XRequestId *string `json:"X-Request-Id,omitempty"`

	// 语言参数，默认为中文zh-CN，英文为en-US。
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// 查询偏移量,若超过最大数量，则返回最后一页。
	Offset *int32 `json:"offset,omitempty"`

	// 查询数量。 默认值：10。
	Limit *int32 `json:"limit,omitempty"`

	// 搜索条件，支持名称、手机、邮箱、帐号、第三方帐号模糊搜索。
	SearchKey *string `json:"searchKey,omitempty"`
}

func (o SearchCorpAdminsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchCorpAdminsRequest struct{}"
	}

	return strings.Join([]string{"SearchCorpAdminsRequest", string(data)}, " ")
}
