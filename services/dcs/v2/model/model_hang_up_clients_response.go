package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HangUpClientsResponse Response Object
type HangUpClientsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o HangUpClientsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HangUpClientsResponse struct{}"
	}

	return strings.Join([]string{"HangUpClientsResponse", string(data)}, " ")
}
