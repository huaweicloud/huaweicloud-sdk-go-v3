package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstancesSessionStatisticsRequest Request Object
type ListInstancesSessionStatisticsRequest struct {

	// 节点ID。
	NodeId string `json:"node_id"`
}

func (o ListInstancesSessionStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesSessionStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ListInstancesSessionStatisticsRequest", string(data)}, " ")
}
