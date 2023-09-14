package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecretEventResponse Response Object
type ShowSecretEventResponse struct {
	Event          *Event `json:"event,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowSecretEventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecretEventResponse struct{}"
	}

	return strings.Join([]string{"ShowSecretEventResponse", string(data)}, " ")
}
