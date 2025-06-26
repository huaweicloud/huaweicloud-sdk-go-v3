package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceStatisticsRequest Request Object
type ListInstanceStatisticsRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`
}

func (o ListInstanceStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceStatisticsRequest", string(data)}, " ")
}
