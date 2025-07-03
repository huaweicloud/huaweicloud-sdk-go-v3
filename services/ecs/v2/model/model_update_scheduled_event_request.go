package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateScheduledEventRequest Request Object
type UpdateScheduledEventRequest struct {

	// resource id
	Id string `json:"id"`

	Body *ScheduledEventUpdateBody `json:"body,omitempty"`
}

func (o UpdateScheduledEventRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScheduledEventRequest struct{}"
	}

	return strings.Join([]string{"UpdateScheduledEventRequest", string(data)}, " ")
}
