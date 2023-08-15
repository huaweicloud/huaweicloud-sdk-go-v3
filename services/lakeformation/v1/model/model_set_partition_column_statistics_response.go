package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetPartitionColumnStatisticsResponse Response Object
type SetPartitionColumnStatisticsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetPartitionColumnStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetPartitionColumnStatisticsResponse struct{}"
	}

	return strings.Join([]string{"SetPartitionColumnStatisticsResponse", string(data)}, " ")
}
