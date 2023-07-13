package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStatisticsResponse Response Object
type ListStatisticsResponse struct {

	// 资源数量信息列表。
	Statistics     *[]Statistic `json:"statistics,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListStatisticsResponse", string(data)}, " ")
}
