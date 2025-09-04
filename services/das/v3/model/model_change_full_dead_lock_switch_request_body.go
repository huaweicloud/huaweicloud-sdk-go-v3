package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChangeFullDeadLockSwitchRequestBody struct {

	// 开关
	SwitchOn bool `json:"switch_on"`

	// 引擎
	EngineType string `json:"engine_type"`
}

func (o ChangeFullDeadLockSwitchRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeFullDeadLockSwitchRequestBody struct{}"
	}

	return strings.Join([]string{"ChangeFullDeadLockSwitchRequestBody", string(data)}, " ")
}
