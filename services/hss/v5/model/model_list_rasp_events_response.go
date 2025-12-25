package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRaspEventsResponse Response Object
type ListRaspEventsResponse struct {

	// **参数解释** 符合所有筛选条件的应用防护事件总数，用于分页计算总页数 **取值范围** 取值0-9223372036854775807
	TotalNum *int64 `json:"total_num,omitempty"`

	// **参数解释** 包含查询到的应用防护事件详细信息，每个元素对应一个防护事件的完整数据 **取值范围** 数组长度0-limit（每页显示个数），元素结构符合RaspProtectHistoryResponseInfo定义，数组为空表示无匹配结果
	DataList       *[]RaspProtectHistoryResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o ListRaspEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRaspEventsResponse struct{}"
	}

	return strings.Join([]string{"ListRaspEventsResponse", string(data)}, " ")
}
