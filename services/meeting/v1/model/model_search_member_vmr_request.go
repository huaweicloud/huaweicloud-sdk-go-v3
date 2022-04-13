package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SearchMemberVmrRequest struct {
	// 请求requestId，用来标识一路请求，用于问题跟踪定位，建议使用uuId，若不携带，则后台自动生成

	XRequestId *string `json:"X-Request-Id,omitempty"`
	// 语言参数，默认为中文zh-CN, 英文为en-US

	AcceptLanguage *string `json:"Accept-Language,omitempty"`
	// 查询偏移量,若超过最大数量，则返回最后一页的数据 默认值：0

	Offset *int32 `json:"offset,omitempty"`
	// 查询数量 默认值：0

	Limit *int32 `json:"limit,omitempty"`
	// 搜索条件。支持云会议室名称、ID及分配的用户、硬终端名称模糊搜索。

	SearchKey *string `json:"searchKey,omitempty"`
	// 查询vmr的类型，为null则查询所有。 false:个人云会议室 true:专用云会议室，不带则查询所有

	SpecialVmr *bool `json:"specialVmr,omitempty"`
}

func (o SearchMemberVmrRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchMemberVmrRequest struct{}"
	}

	return strings.Join([]string{"SearchMemberVmrRequest", string(data)}, " ")
}
