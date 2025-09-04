package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateServerGroupRequest Request Object
type ValidateServerGroupRequest struct {
	Body *ValidateSeverGroupReq `json:"body,omitempty"`
}

func (o ValidateServerGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateServerGroupRequest struct{}"
	}

	return strings.Join([]string{"ValidateServerGroupRequest", string(data)}, " ")
}
