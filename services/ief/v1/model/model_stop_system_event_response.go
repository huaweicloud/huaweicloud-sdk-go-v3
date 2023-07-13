package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopSystemEventResponse Response Object
type StopSystemEventResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StopSystemEventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopSystemEventResponse struct{}"
	}

	return strings.Join([]string{"StopSystemEventResponse", string(data)}, " ")
}
