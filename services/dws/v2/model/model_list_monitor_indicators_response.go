package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMonitorIndicatorsResponse Response Object
type ListMonitorIndicatorsResponse struct {
	Body           *[]IndicatorInfo `json:"body,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListMonitorIndicatorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMonitorIndicatorsResponse struct{}"
	}

	return strings.Join([]string{"ListMonitorIndicatorsResponse", string(data)}, " ")
}
