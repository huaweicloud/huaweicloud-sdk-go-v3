package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyVmNicResponse Response Object
type ModifyVmNicResponse struct {
	PortId         *string `json:"port_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ModifyVmNicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyVmNicResponse struct{}"
	}

	return strings.Join([]string{"ModifyVmNicResponse", string(data)}, " ")
}
