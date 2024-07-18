package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVpnUserResponse Response Object
type DeleteVpnUserResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVpnUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVpnUserResponse struct{}"
	}

	return strings.Join([]string{"DeleteVpnUserResponse", string(data)}, " ")
}
