package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEventResponse Response Object
type CreateEventResponse struct {

	// the id of event
	EventId        *int32 `json:"event_id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateEventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEventResponse struct{}"
	}

	return strings.Join([]string{"CreateEventResponse", string(data)}, " ")
}
