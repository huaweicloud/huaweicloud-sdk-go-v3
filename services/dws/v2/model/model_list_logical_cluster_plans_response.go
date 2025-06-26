package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLogicalClusterPlansResponse Response Object
type ListLogicalClusterPlansResponse struct {

	// **参数解释**： 逻辑集群增删计划对象列表。列表最大返回一般不超过20条。 **取值范围**： 不涉及。
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
