package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetTwoFactorLoginConfigResponse Response Object
type SetTwoFactorLoginConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetTwoFactorLoginConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetTwoFactorLoginConfigResponse struct{}"
	}

	return strings.Join([]string{"SetTwoFactorLoginConfigResponse", string(data)}, " ")
}
