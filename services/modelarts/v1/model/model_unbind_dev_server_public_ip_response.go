package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnbindDevServerPublicIpResponse Response Object
type UnbindDevServerPublicIpResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UnbindDevServerPublicIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnbindDevServerPublicIpResponse struct{}"
	}

	return strings.Join([]string{"UnbindDevServerPublicIpResponse", string(data)}, " ")
}
