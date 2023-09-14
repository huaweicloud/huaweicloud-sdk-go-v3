package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSecretEventRequest Request Object
type CreateSecretEventRequest struct {
	Body *CreateSecretEventRequestBody `json:"body,omitempty"`
}

func (o CreateSecretEventRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecretEventRequest struct{}"
	}

	return strings.Join([]string{"CreateSecretEventRequest", string(data)}, " ")
}
