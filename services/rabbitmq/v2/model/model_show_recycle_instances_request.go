package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRecycleInstancesRequest Request Object
type ShowRecycleInstancesRequest struct {
}

func (o ShowRecycleInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRecycleInstancesRequest struct{}"
	}

	return strings.Join([]string{"ShowRecycleInstancesRequest", string(data)}, " ")
}
