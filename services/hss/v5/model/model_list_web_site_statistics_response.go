package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWebSiteStatisticsResponse Response Object
type ListWebSiteStatisticsResponse struct {

	// **参数解释** Web站点统计信息总数 **取值范围** 最小值0，最大值300000
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释** Web站点统计信息列表 **取值范围** 最小值0，最大值300000
	DataList       *[]WebSiteStatisticsResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o ListWebSiteStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWebSiteStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListWebSiteStatisticsResponse", string(data)}, " ")
}
