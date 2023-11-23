package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchIpGroupRequest Request Object
type SwitchIpGroupRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *SwitchIpGroupRequestBody `json:"body,omitempty"`
}

func (o SwitchIpGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchIpGroupRequest struct{}"
	}

	return strings.Join([]string{"SwitchIpGroupRequest", string(data)}, " ")
}
