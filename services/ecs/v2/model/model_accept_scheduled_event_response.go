package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AcceptScheduledEventResponse Response Object
type AcceptScheduledEventResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AcceptScheduledEventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AcceptScheduledEventResponse struct{}"
	}

	return strings.Join([]string{"AcceptScheduledEventResponse", string(data)}, " ")
}
