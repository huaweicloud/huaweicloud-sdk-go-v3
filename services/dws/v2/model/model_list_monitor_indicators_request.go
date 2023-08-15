package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMonitorIndicatorsRequest Request Object
type ListMonitorIndicatorsRequest struct {
}

func (o ListMonitorIndicatorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMonitorIndicatorsRequest struct{}"
	}

	return strings.Join([]string{"ListMonitorIndicatorsRequest", string(data)}, " ")
}
