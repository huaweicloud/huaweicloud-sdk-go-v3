package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteSchedulesRequest Request Object
type BatchDeleteSchedulesRequest struct {
	Body *BatchDeleteSchedulesRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteSchedulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSchedulesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteSchedulesRequest", string(data)}, " ")
}
