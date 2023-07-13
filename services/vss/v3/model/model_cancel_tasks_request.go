package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelTasksRequest Request Object
type CancelTasksRequest struct {
	Body *CancelTasksRequestBody `json:"body,omitempty"`
}

func (o CancelTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelTasksRequest struct{}"
	}

	return strings.Join([]string{"CancelTasksRequest", string(data)}, " ")
}
