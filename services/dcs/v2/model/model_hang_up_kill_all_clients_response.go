package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HangUpKillAllClientsResponse Response Object
type HangUpKillAllClientsResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o HangUpKillAllClientsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HangUpKillAllClientsResponse struct{}"
	}

	return strings.Join([]string{"HangUpKillAllClientsResponse", string(data)}, " ")
}
