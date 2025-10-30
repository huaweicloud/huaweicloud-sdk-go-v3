package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWebAppAndServiceStatisticsResponse Response Object
type ListWebAppAndServiceStatisticsResponse struct {

	// **参数解释** web应用、web服务、数据库指纹的数量 **取值范围** 最小值0，最大值300000\"
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释** web应用、web服务、数据库指纹统计信息列表 **取值范围** 最小值0，最大值10000
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
