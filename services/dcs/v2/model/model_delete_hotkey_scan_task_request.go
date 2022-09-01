package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteHotkeyScanTaskRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 热key分析任务ID。
	HotkeyId string `json:"hotkey_id" xml:"hotkey_id"`
}

func (o DeleteHotkeyScanTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHotkeyScanTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteHotkeyScanTaskRequest", string(data)}, " ")
}
