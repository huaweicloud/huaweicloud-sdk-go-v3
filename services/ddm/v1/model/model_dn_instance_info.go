package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DnInstanceInfo struct {

	// 实例id。
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名称。
	InstanceName *string `json:"instance_name,omitempty"`

	// ip。
	AccessHost *string `json:"access_host,omitempty"`

	// 端口。
	AccessPort *int32 `json:"access_port,omitempty"`

	// 引擎。
	Engine *string `json:"engine,omitempty"`

	// 引擎版本。
	EngineVersion *string `json:"engine_version,omitempty"`
}

func (o DnInstanceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DnInstanceInfo struct{}"
	}

	return strings.Join([]string{"DnInstanceInfo", string(data)}, " ")
}
