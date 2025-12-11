package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserStatisticsResponse Response Object
type ListUserStatisticsResponse struct {

	// **参数解释**: 账号总数 **取值范围**: 最小值0，最大值10000000
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 账号历史变账户统计信息列表动记录列表 **取值范围**: 最小值0，最大值10000
	DataList       *[]UserStatisticInfoResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o ListUserStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListUserStatisticsResponse", string(data)}, " ")
}
