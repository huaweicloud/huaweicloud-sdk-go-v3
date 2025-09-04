package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetUserProfileResponse Response Object
type ResetUserProfileResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResetUserProfileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetUserProfileResponse struct{}"
	}

	return strings.Join([]string{"ResetUserProfileResponse", string(data)}, " ")
}
