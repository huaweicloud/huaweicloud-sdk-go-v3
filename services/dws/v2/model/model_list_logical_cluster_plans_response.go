package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLogicalClusterPlansResponse Response Object
type ListLogicalClusterPlansResponse struct {

	// 显示逻辑集群增删计划
	Plans          *[]LogicalClusterPlanVo `json:"plans,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListLogicalClusterPlansResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogicalClusterPlansResponse struct{}"
	}

	return strings.Join([]string{"ListLogicalClusterPlansResponse", string(data)}, " ")
}
