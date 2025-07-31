package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWebFrameworkStatisticsResponse Response Object
type ListWebFrameworkStatisticsResponse struct {

	// Web框架统计信息总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// Web框架统计信息列表
	DataList       *[]WebFrameworkStatisticsResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o ListWebFrameworkStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWebFrameworkStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListWebFrameworkStatisticsResponse", string(data)}, " ")
}
