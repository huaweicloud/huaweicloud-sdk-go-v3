package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBizMetricsResponse Response Object
type ListBizMetricsResponse struct {
	Data           *ListBizMetricsResultData `json:"data,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListBizMetricsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBizMetricsResponse struct{}"
	}

	return strings.Join([]string{"ListBizMetricsResponse", string(data)}, " ")
}
