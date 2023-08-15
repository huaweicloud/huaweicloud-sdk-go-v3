package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelBroadcastResponse Response Object
type CancelBroadcastResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CancelBroadcastResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelBroadcastResponse struct{}"
	}

	return strings.Join([]string{"CancelBroadcastResponse", string(data)}, " ")
}
