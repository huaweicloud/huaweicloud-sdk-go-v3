package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAuthModeRequest Request Object
type UpdateAuthModeRequest struct {
	Body *UpdateAuthModeRequestBody `json:"body,omitempty"`
}

func (o UpdateAuthModeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAuthModeRequest struct{}"
	}

	return strings.Join([]string{"UpdateAuthModeRequest", string(data)}, " ")
}
