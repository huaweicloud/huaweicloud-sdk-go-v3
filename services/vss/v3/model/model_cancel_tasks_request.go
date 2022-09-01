package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CancelTasksRequest struct {
	Body *CancelTasksRequestBody `json:"body,omitempty" xml:"body"`
}

func (o CancelTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelTasksRequest struct{}"
	}

	return strings.Join([]string{"CancelTasksRequest", string(data)}, " ")
}
