package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTableColumnStatisticsResponse Response Object
type DeleteTableColumnStatisticsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTableColumnStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTableColumnStatisticsResponse struct{}"
	}

	return strings.Join([]string{"DeleteTableColumnStatisticsResponse", string(data)}, " ")
}
