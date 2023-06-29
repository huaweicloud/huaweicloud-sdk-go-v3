package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStatisticsRequest Request Object
type ListStatisticsRequest struct {
}

func (o ListStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ListStatisticsRequest", string(data)}, " ")
}
