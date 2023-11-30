package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCompoundMetricsResponse Response Object
type ListCompoundMetricsResponse struct {
	Data           *CompoundMetricVoSearchResultData `json:"data,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o ListCompoundMetricsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCompoundMetricsResponse struct{}"
	}

	return strings.Join([]string{"ListCompoundMetricsResponse", string(data)}, " ")
}
