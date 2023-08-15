package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMonitorIndicatorDataResponse Response Object
type ListMonitorIndicatorDataResponse struct {
	Body           *[]TrendQueryDataResp `json:"body,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListMonitorIndicatorDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMonitorIndicatorDataResponse struct{}"
	}

	return strings.Join([]string{"ListMonitorIndicatorDataResponse", string(data)}, " ")
}
