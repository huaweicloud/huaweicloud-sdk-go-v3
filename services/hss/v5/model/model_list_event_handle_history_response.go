package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEventHandleHistoryResponse Response Object
type ListEventHandleHistoryResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// 告警事件历史处置记录
	DataList       *[]EventHandleHistory `json:"data_list,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListEventHandleHistoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventHandleHistoryResponse struct{}"
	}

	return strings.Join([]string{"ListEventHandleHistoryResponse", string(data)}, " ")
}
