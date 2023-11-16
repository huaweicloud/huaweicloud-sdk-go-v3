package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTagForProtectedIpResponse Response Object
type UpdateTagForProtectedIpResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTagForProtectedIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTagForProtectedIpResponse struct{}"
	}

	return strings.Join([]string{"UpdateTagForProtectedIpResponse", string(data)}, " ")
}
