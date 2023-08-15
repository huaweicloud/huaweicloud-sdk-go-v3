package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStatisticsOfRunningInstancesRequest Request Object
type ListStatisticsOfRunningInstancesRequest struct {
}

func (o ListStatisticsOfRunningInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStatisticsOfRunningInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListStatisticsOfRunningInstancesRequest", string(data)}, " ")
}
