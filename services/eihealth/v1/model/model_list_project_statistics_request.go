package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectStatisticsRequest Request Object
type ListProjectStatisticsRequest struct {
}

func (o ListProjectStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ListProjectStatisticsRequest", string(data)}, " ")
}
