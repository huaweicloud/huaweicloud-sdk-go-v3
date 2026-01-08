package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateConfigRequest Request Object
type ValidateConfigRequest struct {
	Body *AuthMethodConfigRequest `json:"body,omitempty"`
}

func (o ValidateConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateConfigRequest struct{}"
	}

	return strings.Join([]string{"ValidateConfigRequest", string(data)}, " ")
}
