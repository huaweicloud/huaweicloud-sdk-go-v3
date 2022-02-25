package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SearchMemberVmrResponse struct {
	// 页面起始页，从0开始

	Offset *int32 `json:"offset,omitempty"`
	// 每页显示的条目数量。 默认值：10。

	Limit *int32 `json:"limit,omitempty"`
	// 总数量。

	Count *int32 `json:"count,omitempty"`
	// 查询到的用户云会议室列表

	Data           *[]QueryVmrResultDto `json:"data,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o SearchMemberVmrResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchMemberVmrResponse struct{}"
	}

	return strings.Join([]string{"SearchMemberVmrResponse", string(data)}, " ")
}
