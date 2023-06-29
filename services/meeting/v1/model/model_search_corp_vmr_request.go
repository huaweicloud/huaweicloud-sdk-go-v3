package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchCorpVmrRequest Request Object
type SearchCorpVmrRequest struct {

	// 请求requestId，用来标识一路请求，用于问题跟踪定位，建议使用UUID，若不携带，则后台自动生成。
	XRequestId *string `json:"X-Request-Id,omitempty"`

	// 语言参数，默认为中文zh-CN，英文为en-US。
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// 查询偏移量,若超过最大数量，则返回最后一页的数据。 默认值：0。
	Offset *int32 `json:"offset,omitempty"`

	// 查询数量。 默认值：10。
	Limit *int32 `json:"limit,omitempty"`

	// 搜索条件。支持云会议室名称、ID及分配的用户、硬终端名称模糊搜索。
	SearchKey *string `json:"searchKey,omitempty"`

	// VMR模式。不填则默认为云会议室。 - 1：云会议室 - 2：网络研讨会
	VmrMode *int32 `json:"vmrMode,omitempty"`

	// 云会议室状态。不填则查询所有。 * 0：正常 * 1：停用 * 2：未分配
	Status *int32 `json:"status,omitempty"`
}

func (o SearchCorpVmrRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchCorpVmrRequest struct{}"
	}

	return strings.Join([]string{"SearchCorpVmrRequest", string(data)}, " ")
}
