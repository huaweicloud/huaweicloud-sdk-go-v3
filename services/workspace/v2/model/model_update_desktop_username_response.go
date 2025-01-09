package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDesktopUsernameResponse Response Object
type UpdateDesktopUsernameResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateDesktopUsernameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDesktopUsernameResponse struct{}"
	}

	return strings.Join([]string{"UpdateDesktopUsernameResponse", string(data)}, " ")
}
