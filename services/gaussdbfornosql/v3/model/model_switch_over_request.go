package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchOverRequest Request Object
type SwitchOverRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o SwitchOverRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchOverRequest struct{}"
	}

	return strings.Join([]string{"SwitchOverRequest", string(data)}, " ")
}
