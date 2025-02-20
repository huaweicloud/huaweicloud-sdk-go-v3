package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateApplicationViewRequest Request Object
type BatchCreateApplicationViewRequest struct {
	Body *BatchCreateApplicationViewRequestBody `json:"body,omitempty"`
}

func (o BatchCreateApplicationViewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateApplicationViewRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateApplicationViewRequest", string(data)}, " ")
}
