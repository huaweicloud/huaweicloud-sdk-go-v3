package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetViewSqlStatisticsRequest Request Object
type ResetViewSqlStatisticsRequest struct {

	// insatnceId
	InstanceId string `json:"instance_id"`
}

func (o ResetViewSqlStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetViewSqlStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ResetViewSqlStatisticsRequest", string(data)}, " ")
}
