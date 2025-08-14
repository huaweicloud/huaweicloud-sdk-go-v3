package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppStatisticsResponse Response Object
type ListAppStatisticsResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 进程统计信息列表 **取值范围**: 最小值0，最大值10000
	DataList       *[]AppStatisticResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListAppStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListAppStatisticsResponse", string(data)}, " ")
}
