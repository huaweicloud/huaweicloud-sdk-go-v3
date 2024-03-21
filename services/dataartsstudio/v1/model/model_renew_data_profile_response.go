package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RenewDataProfileResponse Response Object
type RenewDataProfileResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RenewDataProfileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RenewDataProfileResponse struct{}"
	}

	return strings.Join([]string{"RenewDataProfileResponse", string(data)}, " ")
}
