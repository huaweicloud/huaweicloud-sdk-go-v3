package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHyperinstanceGroupsCapacityRequest Request Object
type ListHyperinstanceGroupsCapacityRequest struct {
	Body *HyperinstanceGroupsCapacityRequest `json:"body,omitempty"`
}

func (o ListHyperinstanceGroupsCapacityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHyperinstanceGroupsCapacityRequest struct{}"
	}

	return strings.Join([]string{"ListHyperinstanceGroupsCapacityRequest", string(data)}, " ")
}
