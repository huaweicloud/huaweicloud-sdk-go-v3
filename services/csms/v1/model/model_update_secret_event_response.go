package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecretEventResponse Response Object
type UpdateSecretEventResponse struct {
	Event          *Event `json:"event,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdateSecretEventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecretEventResponse struct{}"
	}

	return strings.Join([]string{"UpdateSecretEventResponse", string(data)}, " ")
}
