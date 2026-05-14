package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateWhitelistUserRequest Request Object
type ValidateWhitelistUserRequest struct {
	Body *ValidateWhitelistUserRequestBody `json:"body,omitempty"`
}

func (o ValidateWhitelistUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateWhitelistUserRequest struct{}"
	}

	return strings.Join([]string{"ValidateWhitelistUserRequest", string(data)}, " ")
}
