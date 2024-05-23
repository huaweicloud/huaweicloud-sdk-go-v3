package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchInstanceVpcResponse Response Object
type SwitchInstanceVpcResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SwitchInstanceVpcResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchInstanceVpcResponse struct{}"
	}

	return strings.Join([]string{"SwitchInstanceVpcResponse", string(data)}, " ")
}
