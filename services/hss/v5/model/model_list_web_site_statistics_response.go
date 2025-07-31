package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWebSiteStatisticsResponse Response Object
type ListWebSiteStatisticsResponse struct {

	// Web站点统计信息总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// Web站点统计信息列表
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
