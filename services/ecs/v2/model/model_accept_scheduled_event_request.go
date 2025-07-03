package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AcceptScheduledEventRequest Request Object
type AcceptScheduledEventRequest struct {

	// resource id
	Id string `json:"id"`

	Body *ScheduledEventAcceptBody `json:"body,omitempty"`
}

func (o AcceptScheduledEventRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AcceptScheduledEventRequest struct{}"
	}

	return strings.Join([]string{"AcceptScheduledEventRequest", string(data)}, " ")
}
