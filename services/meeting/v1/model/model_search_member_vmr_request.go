package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SearchMemberVmrRequest struct {

	// 请求requestId，用来标识一路请求，用于问题跟踪定位，建议使用UUID，若不携带，则后台自动生成。
	XRequestId *string `json:"X-Request-Id,omitempty"`

	// 语言参数，默认为中文zh-CN，英文为en-US。
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// 查询偏移量,若超过最大数量，则返回最后一页的数据。 默认值：0。
	Offset *int32 `json:"offset,omitempty"`

	// 查询数量。 默认值：10。
	Limit *int32 `json:"limit,omitempty"`

	// 搜索条件。支持云会议室名称、ID模糊搜索。
	SearchKey *string `json:"searchKey,omitempty"`

	// 查询VMR的类型。不填则查询所有类型。 * false:个人会议ID * true:云会议室
	SpecialVmr *bool `json:"specialVmr,omitempty"`
}

func (o SearchMemberVmrRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchMemberVmrRequest struct{}"
	}

	return strings.Join([]string{"SearchMemberVmrRequest", string(data)}, " ")
}
