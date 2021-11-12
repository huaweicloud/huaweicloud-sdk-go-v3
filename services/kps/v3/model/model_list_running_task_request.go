package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListRunningTaskRequest struct {
}

func (o ListRunningTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRunningTaskRequest struct{}"
	}

	return strings.Join([]string{"ListRunningTaskRequest", string(data)}, " ")
}
