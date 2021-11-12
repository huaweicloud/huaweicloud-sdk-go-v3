package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListFailedTaskRequest struct {
}

func (o ListFailedTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFailedTaskRequest struct{}"
	}

	return strings.Join([]string{"ListFailedTaskRequest", string(data)}, " ")
}
