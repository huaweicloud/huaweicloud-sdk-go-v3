package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTopnTrafficStatisticsResponse Response Object
type ListTopnTrafficStatisticsResponse struct {

	// 响应结果
	Results        *[]ResultsTopnBody `json:"results,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListTopnTrafficStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopnTrafficStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListTopnTrafficStatisticsResponse", string(data)}, " ")
}
