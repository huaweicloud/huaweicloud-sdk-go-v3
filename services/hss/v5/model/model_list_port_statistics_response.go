package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPortStatisticsResponse Response Object
type ListPortStatisticsResponse struct {

	// **参数解释**: 开放端口总数 **取值范围**: 最小值0， 最大值10000
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 开放端口统计信息列表 **取值范围**: 最小值0， 最大值10000
	DataList       *[]PortStatisticResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListPortStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPortStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListPortStatisticsResponse", string(data)}, " ")
}
