package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchStartServerResponse Response Object
type BatchStartServerResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchStartServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStartServerResponse struct{}"
	}

	return strings.Join([]string{"BatchStartServerResponse", string(data)}, " ")
}
