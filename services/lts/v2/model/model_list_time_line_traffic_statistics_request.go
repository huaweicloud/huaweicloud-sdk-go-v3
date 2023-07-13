package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTimeLineTrafficStatisticsRequest Request Object
type ListTimeLineTrafficStatisticsRequest struct {

	// 时区
	Timezone string `json:"timezone"`

	Body *TimelineTrafficStatisticsRequestBody `json:"body,omitempty"`
}

func (o ListTimeLineTrafficStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTimeLineTrafficStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ListTimeLineTrafficStatisticsRequest", string(data)}, " ")
}
