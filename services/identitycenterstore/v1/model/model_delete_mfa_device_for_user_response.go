package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMfaDeviceForUserResponse Response Object
type DeleteMfaDeviceForUserResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteMfaDeviceForUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMfaDeviceForUserResponse struct{}"
	}

	return strings.Join([]string{"DeleteMfaDeviceForUserResponse", string(data)}, " ")
}
