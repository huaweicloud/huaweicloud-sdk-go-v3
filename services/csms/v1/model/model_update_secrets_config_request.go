package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecretsConfigRequest Request Object
type UpdateSecretsConfigRequest struct {
	Body *UpdateSecretsConfigRequestBody `json:"body,omitempty"`
}

func (o UpdateSecretsConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecretsConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateSecretsConfigRequest", string(data)}, " ")
}
