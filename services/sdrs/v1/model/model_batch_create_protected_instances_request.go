package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateProtectedInstancesRequest Request Object
type BatchCreateProtectedInstancesRequest struct {
	Body *BatchCreateProtectedInstancesRequestBody `json:"body,omitempty"`
}

func (o BatchCreateProtectedInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateProtectedInstancesRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateProtectedInstancesRequest", string(data)}, " ")
}
