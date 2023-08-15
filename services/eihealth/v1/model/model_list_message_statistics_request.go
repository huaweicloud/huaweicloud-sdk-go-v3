package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMessageStatisticsRequest Request Object
type ListMessageStatisticsRequest struct {
}

func (o ListMessageStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMessageStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ListMessageStatisticsRequest", string(data)}, " ")
}
