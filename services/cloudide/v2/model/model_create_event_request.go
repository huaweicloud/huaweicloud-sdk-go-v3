package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEventRequest Request Object
type CreateEventRequest struct {
	Body *EventSchema `json:"body,omitempty"`
}

func (o CreateEventRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEventRequest struct{}"
	}

	return strings.Join([]string{"CreateEventRequest", string(data)}, " ")
}
