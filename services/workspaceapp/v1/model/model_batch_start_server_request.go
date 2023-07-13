package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchStartServerRequest Request Object
type BatchStartServerRequest struct {
	Body *BatchStartServerReq `json:"body,omitempty"`
}

func (o BatchStartServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStartServerRequest struct{}"
	}

	return strings.Join([]string{"BatchStartServerRequest", string(data)}, " ")
}
