package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchIpGroupResponse Response Object
type SwitchIpGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SwitchIpGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchIpGroupResponse struct{}"
	}

	return strings.Join([]string{"SwitchIpGroupResponse", string(data)}, " ")
}
