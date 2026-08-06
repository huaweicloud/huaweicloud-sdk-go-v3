package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchFullsqlSwitchRequest Request Object
type SwitchFullsqlSwitchRequest struct {

	// 引擎类型
	EngineType *string `json:"engine_type,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`
}

func (o SwitchFullsqlSwitchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchFullsqlSwitchRequest struct{}"
	}

	return strings.Join([]string{"SwitchFullsqlSwitchRequest", string(data)}, " ")
}
