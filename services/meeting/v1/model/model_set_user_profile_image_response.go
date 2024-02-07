package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetUserProfileImageResponse Response Object
type SetUserProfileImageResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetUserProfileImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetUserProfileImageResponse struct{}"
	}

	return strings.Join([]string{"SetUserProfileImageResponse", string(data)}, " ")
}
