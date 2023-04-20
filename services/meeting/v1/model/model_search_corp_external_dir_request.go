package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SearchCorpExternalDirRequest struct {

	// 请求requestId，用来标识一路请求，用于问题跟踪定位，建议使用UUID，若不携带，则后台自动生成。
	XRequestId *string `json:"X-Request-Id,omitempty"`

	// 语言参数，默认为中文zh-CN，英文为en-US。
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// 查询偏移量,若超过最大数量，则返回最后一页的数据。 默认值：0。
	Offset *int32 `json:"offset,omitempty"`

	// 查询数量。 默认值：10。
	Limit *int32 `json:"limit,omitempty"`

	// 搜索条件。支持名称、手机、邮箱模糊搜索。
	SearchKey *string `json:"searchKey,omitempty"`

	// 搜索范围。默认值为ALL。 * PERSONAL - 查询个人外部联系人 * CORP - 查询企业外部联系人 * ALL - 查询所有外部联系人
	SearchScope *string `json:"searchScope,omitempty"`
}

func (o SearchCorpExternalDirRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchCorpExternalDirRequest struct{}"
	}

	return strings.Join([]string{"SearchCorpExternalDirRequest", string(data)}, " ")
}
