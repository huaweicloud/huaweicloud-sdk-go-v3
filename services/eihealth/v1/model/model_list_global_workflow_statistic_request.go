package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGlobalWorkflowStatisticRequest Request Object
type ListGlobalWorkflowStatisticRequest struct {
}

func (o ListGlobalWorkflowStatisticRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGlobalWorkflowStatisticRequest struct{}"
	}

	return strings.Join([]string{"ListGlobalWorkflowStatisticRequest", string(data)}, " ")
}
