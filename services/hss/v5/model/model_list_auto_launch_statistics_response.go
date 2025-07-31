package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAutoLaunchStatisticsResponse Response Object
type ListAutoLaunchStatisticsResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**： 自启动项统计信息列表 **取值范围**： 不涉及
	DataList       *[]AutoLaunchStatisticsResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o ListAutoLaunchStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAutoLaunchStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListAutoLaunchStatisticsResponse", string(data)}, " ")
}
