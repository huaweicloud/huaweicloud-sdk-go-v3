package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetHarmonySoftBusKeyResponse Response Object
type ResetHarmonySoftBusKeyResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResetHarmonySoftBusKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetHarmonySoftBusKeyResponse struct{}"
	}

	return strings.Join([]string{"ResetHarmonySoftBusKeyResponse", string(data)}, " ")
}
