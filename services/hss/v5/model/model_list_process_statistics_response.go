package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProcessStatisticsResponse Response Object
type ListProcessStatisticsResponse struct {

	// **参数解释** : 进程统计信息总数 **约束限制** : 不涉及 **取值范围** : 最小值0，最大值10000 **默认取值** : 不涉及
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**： 进程统计信息列表 **约束限制** : 不涉及 **取值范围**： 不涉及 **默认取值** : 不涉及
	DataList       *[]ProcessStatisticResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ListProcessStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProcessStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListProcessStatisticsResponse", string(data)}, " ")
}
