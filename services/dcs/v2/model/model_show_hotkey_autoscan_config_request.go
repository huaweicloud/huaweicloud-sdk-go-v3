package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHotkeyAutoscanConfigRequest Request Object
type ShowHotkeyAutoscanConfigRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ShowHotkeyAutoscanConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHotkeyAutoscanConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowHotkeyAutoscanConfigRequest", string(data)}, " ")
}
