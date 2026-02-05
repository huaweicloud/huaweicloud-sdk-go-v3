package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchIpGroupResponse Response Object
type SwitchIpGroupResponse struct {

	// **参数解释**：  是否开启弹性负载均衡ip组。  **参数范围**：  不涉及。
	EnableIpgroup  *bool `json:"enable_ipgroup,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o SwitchIpGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchIpGroupResponse struct{}"
	}

	return strings.Join([]string{"SwitchIpGroupResponse", string(data)}, " ")
}
