package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHistoryPeriodResponse Response Object
type ListHistoryPeriodResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 历史报告统计周期列表 **取值范围**: 取值0-2000个HistoryPeriodInfo对象
	DataList       *[]HistoryPeriodInfo `json:"data_list,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListHistoryPeriodResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHistoryPeriodResponse struct{}"
	}

	return strings.Join([]string{"ListHistoryPeriodResponse", string(data)}, " ")
}
