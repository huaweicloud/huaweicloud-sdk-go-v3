package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWebAppAndServiceStatisticsResponse Response Object
type ListWebAppAndServiceStatisticsResponse struct {

	// 不同的数据库资产指纹的数量
	TotalNum *int32 `json:"total_num,omitempty"`

	// 数据库资产指纹统计信息列表
	DataList       *[]WebAppAndServiceStatisticResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                                      `json:"-"`
}

func (o ListWebAppAndServiceStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWebAppAndServiceStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListWebAppAndServiceStatisticsResponse", string(data)}, " ")
}
