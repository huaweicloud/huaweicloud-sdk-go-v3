package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchStartInstanceRequest struct {
	Body *BatchStartInstanceRequestBody `json:"body,omitempty"`
}

func (o BatchStartInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStartInstanceRequest struct{}"
	}

	return strings.Join([]string{"BatchStartInstanceRequest", string(data)}, " ")
}
