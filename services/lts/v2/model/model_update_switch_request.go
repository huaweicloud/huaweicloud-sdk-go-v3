package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSwitchRequest Request Object
type UpdateSwitchRequest struct {

	// 开关参数
	LogConvergeSwitch string `json:"log_converge_switch"`
}

func (o UpdateSwitchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSwitchRequest struct{}"
	}

	return strings.Join([]string{"UpdateSwitchRequest", string(data)}, " ")
}
