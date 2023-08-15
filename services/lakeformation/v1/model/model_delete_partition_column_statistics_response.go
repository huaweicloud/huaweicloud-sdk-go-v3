package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePartitionColumnStatisticsResponse Response Object
type DeletePartitionColumnStatisticsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePartitionColumnStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePartitionColumnStatisticsResponse struct{}"
	}

	return strings.Join([]string{"DeletePartitionColumnStatisticsResponse", string(data)}, " ")
}
