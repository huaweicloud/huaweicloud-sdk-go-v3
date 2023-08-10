package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PutEventsResponse Response Object
type PutEventsResponse struct {
	Body *string `json:"body,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o PutEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutEventsResponse struct{}"
	}

	return strings.Join([]string{"PutEventsResponse", string(data)}, " ")
}
