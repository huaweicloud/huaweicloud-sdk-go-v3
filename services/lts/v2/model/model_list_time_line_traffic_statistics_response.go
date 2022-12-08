package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListTimeLineTrafficStatisticsResponse struct {

	// 响应结果
	Results        *[]Resulits `json:"results,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListTimeLineTrafficStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTimeLineTrafficStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListTimeLineTrafficStatisticsResponse", string(data)}, " ")
}
