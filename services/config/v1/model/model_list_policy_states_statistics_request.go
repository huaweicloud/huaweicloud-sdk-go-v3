package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPolicyStatesStatisticsRequest Request Object
type ListPolicyStatesStatisticsRequest struct {
}

func (o ListPolicyStatesStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPolicyStatesStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ListPolicyStatesStatisticsRequest", string(data)}, " ")
}
