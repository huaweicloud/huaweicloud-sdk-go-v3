package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstancesStatusRequest Request Object
type ListInstancesStatusRequest struct {
}

func (o ListInstancesStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesStatusRequest struct{}"
	}

	return strings.Join([]string{"ListInstancesStatusRequest", string(data)}, " ")
}
