package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateHotkeyAutoScanConfigRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	Body *AutoscanConfigRequest `json:"body,omitempty" xml:"body"`
}

func (o UpdateHotkeyAutoScanConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHotkeyAutoScanConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateHotkeyAutoScanConfigRequest", string(data)}, " ")
}
