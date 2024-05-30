package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBizMetricDimensionsResponse Response Object
type ListBizMetricDimensionsResponse struct {
	Data           *ListBizMetricDimensionsResultData `json:"data,omitempty"`
	HttpStatusCode int                                `json:"-"`
}

func (o ListBizMetricDimensionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBizMetricDimensionsResponse struct{}"
	}

	return strings.Join([]string{"ListBizMetricDimensionsResponse", string(data)}, " ")
}
