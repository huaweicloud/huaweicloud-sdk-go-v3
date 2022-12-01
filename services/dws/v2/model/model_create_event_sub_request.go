package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateEventSubRequest struct {
	Body *EventSubRequest `json:"body,omitempty"`
}

func (o CreateEventSubRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEventSubRequest struct{}"
	}

	return strings.Join([]string{"CreateEventSubRequest", string(data)}, " ")
}
