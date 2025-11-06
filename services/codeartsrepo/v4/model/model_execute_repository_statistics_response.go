package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteRepositoryStatisticsResponse Response Object
type ExecuteRepositoryStatisticsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExecuteRepositoryStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteRepositoryStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ExecuteRepositoryStatisticsResponse", string(data)}, " ")
}
