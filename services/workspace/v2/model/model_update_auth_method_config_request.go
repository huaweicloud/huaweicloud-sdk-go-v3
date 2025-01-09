package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAuthMethodConfigRequest Request Object
type UpdateAuthMethodConfigRequest struct {
	Body *AuthMethodConfigRequest `json:"body,omitempty"`
}

func (o UpdateAuthMethodConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAuthMethodConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateAuthMethodConfigRequest", string(data)}, " ")
}
