package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSecretEventResponse Response Object
type CreateSecretEventResponse struct {
	Event          *Event `json:"event,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateSecretEventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecretEventResponse struct{}"
	}

	return strings.Join([]string{"CreateSecretEventResponse", string(data)}, " ")
}
