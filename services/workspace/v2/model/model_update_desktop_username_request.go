package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDesktopUsernameRequest Request Object
type UpdateDesktopUsernameRequest struct {
	Body *UpdateDesktopUsernameReq `json:"body,omitempty"`
}

func (o UpdateDesktopUsernameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDesktopUsernameRequest struct{}"
	}

	return strings.Join([]string{"UpdateDesktopUsernameRequest", string(data)}, " ")
}
