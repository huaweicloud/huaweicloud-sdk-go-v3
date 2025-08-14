package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMfaDeviceForUserResponse Response Object
type UpdateMfaDeviceForUserResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateMfaDeviceForUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMfaDeviceForUserResponse struct{}"
	}

	return strings.Join([]string{"UpdateMfaDeviceForUserResponse", string(data)}, " ")
}
