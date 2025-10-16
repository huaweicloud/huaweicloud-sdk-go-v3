package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPolicyStatesStatisticsResponse Response Object
type ListPolicyStatesStatisticsResponse struct {

	// 合规统计结果
	Value          *[]PolicyStatesStatistics `json:"value,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListPolicyStatesStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPolicyStatesStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListPolicyStatesStatisticsResponse", string(data)}, " ")
}
