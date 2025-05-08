package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceConsumerGroupResponse Response Object
type ListInstanceConsumerGroupResponse struct {
	Group          *DescribeGroupsRespGroup `json:"group,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListInstanceConsumerGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceConsumerGroupResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceConsumerGroupResponse", string(data)}, " ")
}
