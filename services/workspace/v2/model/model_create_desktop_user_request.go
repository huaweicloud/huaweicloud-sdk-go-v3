package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDesktopUserRequest Request Object
type CreateDesktopUserRequest struct {
	Body *CreateUserRequest `json:"body,omitempty"`
}

func (o CreateDesktopUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDesktopUserRequest struct{}"
	}

	return strings.Join([]string{"CreateDesktopUserRequest", string(data)}, " ")
}
