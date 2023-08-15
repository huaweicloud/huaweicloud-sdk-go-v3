package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUnusedDesktopsRequest Request Object
type ListUnusedDesktopsRequest struct {

	// 从查询结果中的第几条数据开始返回,用于分页查询，取值范围0-2000，默认从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// 查询结果中想要返回的信息条目数量,用于分页查询，取值范围0-2000，默认值1000。
	Limit *int32 `json:"limit,omitempty"`

	// 开始时间：由日期加时间组成，UTC格式，格式：yyyy-MM-ddTHH:mm:ss.SSSZ，若未输入，则查询现在到前一天的未使用的桌面。
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间：由日期加时间组成，UTC格式，格式：yyyy-MM-ddTHH:mm:ss.SSSZ，若未输入，则查询现在到前一天的未使用的桌面。
	EndTime *string `json:"end_time,omitempty"`
}

func (o ListUnusedDesktopsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUnusedDesktopsRequest struct{}"
	}

	return strings.Join([]string{"ListUnusedDesktopsRequest", string(data)}, " ")
}
