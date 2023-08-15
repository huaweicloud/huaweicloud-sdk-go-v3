package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PushEventsResponse Response Object
type PushEventsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o PushEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PushEventsResponse struct{}"
	}

	return strings.Join([]string{"PushEventsResponse", string(data)}, " ")
}
