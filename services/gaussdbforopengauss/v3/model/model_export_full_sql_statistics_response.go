package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportFullSqlStatisticsResponse Response Object
type ExportFullSqlStatisticsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExportFullSqlStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportFullSqlStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ExportFullSqlStatisticsResponse", string(data)}, " ")
}
