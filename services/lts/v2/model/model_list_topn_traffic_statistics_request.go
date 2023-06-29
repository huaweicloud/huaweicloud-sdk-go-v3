package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTopnTrafficStatisticsRequest Request Object
type ListTopnTrafficStatisticsRequest struct {
	Body *TopnRequstBody `json:"body,omitempty"`
}

func (o ListTopnTrafficStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopnTrafficStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ListTopnTrafficStatisticsRequest", string(data)}, " ")
}
