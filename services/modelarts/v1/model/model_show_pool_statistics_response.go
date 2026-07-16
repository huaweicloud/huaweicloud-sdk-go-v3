package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPoolStatisticsResponse Response Object
type ShowPoolStatisticsResponse struct {
	Statistics *PoolStatisticsStatistics `json:"statistics,omitempty"`

	// **参数描述**： 统计的时间。 **取值范围**： 不涉及。
	OperationTime  *string `json:"operationTime,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPoolStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPoolStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowPoolStatisticsResponse", string(data)}, " ")
}
