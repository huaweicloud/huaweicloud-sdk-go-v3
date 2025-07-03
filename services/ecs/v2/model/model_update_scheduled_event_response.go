package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateScheduledEventResponse Response Object
type UpdateScheduledEventResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateScheduledEventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScheduledEventResponse struct{}"
	}

	return strings.Join([]string{"UpdateScheduledEventResponse", string(data)}, " ")
}
