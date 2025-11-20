package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AvailableDnInstance struct {

	// 实例id。
	Id *string `json:"id,omitempty"`

	// 实例状态。
	Status *string `json:"status,omitempty"`

	// 实例名称。
	Name *string `json:"name,omitempty"`

	// 引擎名称。
	EngineName *string `json:"engine_name,omitempty"`

	// 引擎版本。
	EngineVersion *string `json:"engine_version,omitempty"`

	// 私有ip。
	PrivateIp *string `json:"private_ip,omitempty"`

	// 端口。
	Port *int32 `json:"port,omitempty"`

	// 时区。
	TimeZone *string `json:"time_zone,omitempty"`
}

func (o AvailableDnInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AvailableDnInstance struct{}"
	}

	return strings.Join([]string{"AvailableDnInstance", string(data)}, " ")
}
